from sedna.datasources import BaseDataSource
from sedna.common.class_factory import ClassFactory, ClassType

from .base_task_allocation import BaseTaskAllocation

@ClassFactory.register(ClassType.STP)
class TaskAllocationByOrigin(BaseTaskAllocation):
    """
    Corresponding to `TaskDefinitionByOrigin`

    Parameters
    ----------
    task_extractor : Dict
        used to predict target tasks for each inference samples
    origins: List[Metadata]
        metadata is usually a class feature
        label with a finite values.
    """

    def __init__(self, task_extractor, **kwargs):
        super(TaskAllocationByOrigin, self).__init__(task_extractor)
        self.default_origin = kwargs.get("default", None)

    def __call__(self, samples: BaseDataSource):
        if self.default_origin:
            return samples, [int(self.task_extractor.get(
                self.default_origin))] * len(samples.x)

        cities = [
            "aachen",
            "berlin",
            "bochum",
            "bremen",
            "cologne",
            "darmstadt",
            "dusseldorf",
            "erfurt",
            "hamburg",
            "hanover",
            "jena",
            "krefeld",
            "monchengladbach",
            "strasbourg",
            "stuttgart",
            "tubingen",
            "ulm",
            "weimar",
            "zurich"]

        sample_origins = []
        for _x in samples.x:
            is_real = False
            for city in cities:
                if city in _x[0]:
                    is_real = True
                    sample_origins.append("real")
                    break
            if not is_real:
                sample_origins.append("sim")

        allocations = [int(self.task_extractor.get(sample_origin))
                       for sample_origin in sample_origins]

        return samples, allocations
