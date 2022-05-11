# Copyright 2021 The KubeEdge Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from abc import ABC, abstractmethod
from os.path import isfile, join

import os
import pickle
import queue
import threading
import time
import traceback
import uuid
from typing import List
from sedna.common.log import LOGGER
from sedna.core.multi_edge_tracking.plugins import PLUGIN, PluggableModel, PluggableNetworkService
from sedna.core.multi_edge_tracking.utils import get_parameters
from sedna.datasources.kafka.kafka_manager import KafkaConsumerThread, KafkaProducer
from distutils import util
from collections import deque

class BaseService(ABC):
    """
    Base wrapper for video analytics, feature extraction, and reid services
    """

    def __init__(self, consumer_topics = [], producer_topics = [], plugins : List[PluggableNetworkService] = [], models : List[PluggableModel] = [], timeout = 10, asynchronous = False):

        LOGGER.info(f"Loaded plugins for this wrapper: {list(map(lambda x: x.kind, plugins))}")
        self.plugins_kind = list(map(lambda x: x.kind, plugins))
        self.plugins = plugins
        self.asynchronous = asynchronous
        self.models = models

        if len(self.models) > 1:
            LOGGER.info("Provided multiple AI executors")
            assert all(isinstance(x, type(models[0])) for x in self.models), "AI executors mixin is not supported!"

        self.batch_size = int(get_parameters('batch_size', 1))

        if self.asynchronous:
            LOGGER.info("Create queue for asynchronous processing.")
            self.sync_queue = deque()
        else:
            LOGGER.info("Create queue for synchronous processing.")
            self.accumulator = queue.Queue(maxsize=self.batch_size)

        # This variables are used to control the data ingestion rate when processing a video.
        self.ingestion_rate = 0
        self.processing_rate = 1
        self.last_put = 0
        self.last_fetch = 0

        self.timeout = timeout

        self._init_kafka_connection(consumer_topics, producer_topics)

        self._post_init()

    def _init_kafka_connection(self, consumer_topics , producer_topics):
        self.kafka_enabled = bool(util.strtobool(get_parameters("KAFKA_ENABLED", "False")))

        if self.kafka_enabled:
            LOGGER.debug("Kafka support enabled in YAML file")
            self.kafka_address = get_parameters("KAFKA_BIND_IPS", [])
            self.kafka_port = get_parameters("KAFKA_BIND_PORTS", [])

            if isinstance(self.kafka_address, str):
                LOGGER.debug(f"Parsing string received from K8s controller {self.kafka_address},{self.kafka_port}")
                self.kafka_address = self.kafka_address.split("|")
                self.kafka_port = self.kafka_port.split("|")
            
            if producer_topics:
                self.producer = KafkaProducer(self.kafka_address, self.kafka_port, topic=producer_topics, asynchronous=self.asynchronous)
            if consumer_topics:
                self.consumer = KafkaConsumerThread(self.kafka_address, self.kafka_port, topic=consumer_topics, callback=self.put)

            LOGGER.info(f"Connection to Kafka broker/s {self.kafka_address}{self.kafka_port} completed.")
            LOGGER.info(f"Consumer topics are {consumer_topics}.")
            LOGGER.info(f"Producer topics are {producer_topics}.")
        
        return


    def _post_init(self):
        threading.Thread(target=self.fetch_data, daemon=True).start()
        return

    # Use asynch mode for ingesting a stream (e.g., RTSP).
    # Use synch mode when reading from disk (e.g., a video file).
    def put(self, data):
        data = self.preprocess(data)
        if data:
            return self._put_data_asynchronous(data) if self.asynchronous else self._put_data_synchronous(data)

    # WARNING: The data sent to the process_data function is always flattened.
    # This means that a list[list[list] will always be transformed into a flat list.
    def fetch_data(self):
        if self.asynchronous:
            self._fetch_asynchronous()
        else:
            self._fetch_synchronous()

    def _fetch_synchronous(self):
        LOGGER.info("Start synchronous fetch loop.")
        while True:
            if self.accumulator.full():
                self._extract_wrapper_sync(self.batch_size)
            elif self.accumulator.qsize() > 0 and (time.time() - self.last_fetch > self.timeout):
                LOGGER.info("Timeout reached. Processing and flushing the remaining elements of the queue.")
                self._extract_wrapper_sync(self.accumulator.qsize())
            else:
                time.sleep(0.01)

    def _fetch_asynchronous(self):
        LOGGER.info("Start asynchronous fetch loop.")
        while True:
            total_stored_elements = len(self.sync_queue)
            if total_stored_elements >= (self.batch_size):
                try:
                    self._extract_wrapper_async(self.batch_size)
                except Exception as e:
                    LOGGER.error(f"Error processing received data: {e}")
                    traceback.print_exc()
            #if we don't receive data for n seconds, flush the queue
            elif total_stored_elements > 0 and (time.time() - self.last_fetch > self.timeout):
                LOGGER.info("Timeout reached. Processing and flushing the remaining elements of the queue.")
                self._extract_wrapper_async(total_stored_elements)
            else:
                time.sleep(0.01)

    def _extract_wrapper_async(self, amount):
        token = [self.sync_queue.popleft() for _ in range(amount)]
        self.last_fetch = time.time()                       
        self.distribute_data(self.flatten(token))
        self.processing_rate = amount/(time.time()-self.last_fetch)
        LOGGER.debug(f"Data Processing Speed: {self.processing_rate} objects/s")

    def _extract_wrapper_sync(self, amount):                   
        token = [self.accumulator.get() for _ in range(amount)]
        self.last_fetch = time.time()    
        self.distribute_data(self.flatten(token))
        [self.accumulator.task_done() for _ in range(amount)]
        self.accumulator.join()

    def _put_data_asynchronous(self, data):
        self.ingestion_rate = len(data)/(time.time()-self.last_put)
        self.sync_queue.append(data)
        LOGGER.debug(f"Data Ingestion Speed: {self.ingestion_rate} objects/s")
        self.last_put = time.time()
        return

    def _put_data_synchronous(self, data):
        self.accumulator.put(data)

        if self.accumulator.full():
            token = [self.accumulator.get() for _ in range(self.batch_size)]
            self.distribute_data(self.flatten(token))
            [self.accumulator.task_done() for _ in range(self.batch_size)]

            self.accumulator.join()
        
        return
        
    def get_plugin(self, plugin_key : PLUGIN):
        try:
            ls = list(filter(lambda n: n.kind == plugin_key.name, self.plugins))[0]
        except IndexError as ie:
            return None
        return ls

    def flatten(self, S):
        if S == []:
            return S
        if isinstance(S[0], list):
            return self.flatten(S[0]) + self.flatten(S[1:])
        return S[:1] + self.flatten(S[1:])

    # Distributes the data in the queue to the models associated to this service
    def distribute_data(self, data = [], **kwargs):
        for ai in self.models:
            self.process_data(ai, data)
        return

    @abstractmethod
    def process_data(self, ai, data, **kwargs):
        return
    
    @abstractmethod
    def update_operational_mode(self, status):
        return

    def preprocess(self, data, **kwargs):
        return data

class FileOperations:
    def read_from_disk(self, path):
        data = []
        try:
            with open(path, 'rb') as diskdata:
                data = pickle.load(diskdata)
        except Exception as ex:
            LOGGER.error(f"Unable to read or load the file! {ex}")

        return data

    def delete_from_disk(self, filename):
        if os.path.exists(filename):
            os.remove(filename)
        else:
            LOGGER.error("The file does not exist.")


    def write_to_disk(self, data, folder, exts=".dat"):
        filename = str(uuid.uuid1())
        with open(f"{folder}{filename}{exts}", 'ab') as result:
            pickle.dump(data, result)


    def get_files_list(self, folder):
        return [join(folder, f) for f in os.listdir(folder) if isfile(join(folder, f))]