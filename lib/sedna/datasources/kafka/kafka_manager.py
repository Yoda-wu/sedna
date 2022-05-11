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

from threading import Thread

from sedna.datasources.kafka.consumer import Consumer
from sedna.datasources.kafka.producer import Producer

class KafkaProducer:
    def __init__(self, address, port, topic=[], asynchronous = False):
        self.producer = Producer(address=address, port=port)
        self.topic = topic
        self.asynchronous = asynchronous

    def write_result(self, data):
        return self.producer.publish_data_asynchronous(data, topic=self.topic) if self.asynchronous else self.producer.publish_data_synchronous(data, topic=self.topic)

class KafkaConsumerThread(Thread):
    def __init__(self, address, port, topic=[], callback = None):
        super().__init__()
        self.consumer = Consumer(address=address, port=port)
        self.callback = callback
        self.topic = topic

        # We do this before actually reading from the topic
        self.consumer.subscribe(self.topic)

        self.daemon = True
        self.start()

    def run(self):
        while not self.consumer.disconnected:
            data = self.consumer.consume_messages_poll()
            if data:
                self.callback(data)