from collections import deque
from json import dump, load
import os
import sys
from concurrent import futures
import time
import grpc
from grpc_reflection.v1alpha import reflection
import threading
from uuid import uuid4

project_root = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
if project_root not in sys.path:
    sys.path.append(project_root)

import proto.fog_pb2 as fog_pb2  # noqa
import proto.fog_pb2_grpc as fog_pb2_grpc  # noqa


class Task:
    def __init__(self, value, sensor_type, task_id=None, result=None, processed=False):
        if task_id is None:
            task_id = str(uuid4())
        self.id = task_id
        self.value = value
        self.sensor_type = sensor_type
        self.result = result
        self.processed = processed

    def to_dict(self):
        return {
            "id": str(self.id),
            "value": self.value,
            "sensor_type": self.sensor_type,
            "result": self.result,
            "processed": self.processed,
        }

    @classmethod
    def from_dict(cls, data):
        return cls(
            value=data["value"],
            sensor_type=data["sensor_type"],
            task_id=data["id"],
            result=data["result"],
            processed=data["processed"],
        )

    def __str__(self):
        return f"Task_id: {self.id}, processed: {self.processed}"


class TaskQueue:
    def __init__(self):
        self.tasks = deque()
        self.lock = threading.Lock()
        self.persistence_file_path = os.path.join(
            os.path.dirname(__file__), "queue_replica.json"
        )
        self.init_persistence_file()
        self.load_tasks_from_replica_queue()

    def init_persistence_file(self):
        if not os.path.exists(self.persistence_file_path):
            with open(self.persistence_file_path, "w") as f:
                print("creating replica file")
                dump([], f)

    def load_tasks_from_replica_queue(self):
        if os.path.exists(self.persistence_file_path):
            with open(self.persistence_file_path, "r") as f:
                backup_data = load(f)
                print(f"loading {len(backup_data)} tasks from replica")
                for task in backup_data:
                    instance_task = Task.from_dict(task)
                    self.add_task(task=instance_task)
                print(f"loading {len(backup_data)} tasks from replica")
        else:
            print("replica file not found!")

    def write_tasks_to_replica_queue(self):
        if os.path.exists(self.persistence_file_path):
            with open(self.persistence_file_path, "w") as f:
                task_list = [task.to_dict() for task in self.tasks]
                dump(task_list, f)
        else:
            print("replica file not found!")

    def add_task(self, task: Task):
        with self.lock:
            self.tasks.append(task)
            self.write_tasks_to_replica_queue()

    def add_task_to_front(self, task: Task):
        with self.lock:
            self.tasks.appendleft(task)
            self.write_tasks_to_replica_queue()

    def update_task(self, task: Task, field, value):
        if hasattr(task, field):
            setattr(task, field, value)
        self.write_tasks_to_replica_queue()
            
    def get_task(self) -> Task:
        with self.lock:
            if self.tasks:
                task = self.tasks.popleft()
                self.write_tasks_to_replica_queue()
                return task
            else:
                return None

    def peek_task(self) -> Task:
        with self.lock:
            if self.tasks:
                return self.tasks[0]
            else:
                return None
        
    def get_all_processed_tasks(self):
        with self.lock:
            processed_tasks = []
            unprocessed_tasks = deque()

            while self.tasks:
                task = self.tasks.popleft()
                if task.processed == True:
                    processed_tasks.append(task)
                else:
                    unprocessed_tasks.append(task)

            self.tasks = unprocessed_tasks
            self.write_tasks_to_replica_queue()
            return processed_tasks


class CloudService(fog_pb2_grpc.CloudServiceServicer):
    def __init__(self):
        self.task_queue = TaskQueue()
        self.edge_ip = os.getenv("EDGE_IP", "localhost")
        self.edge_port = int(os.getenv("EDGE_PORT", 50052))

    def ProcessDataStream(self, request_iterator, context):
        print("Processing data stream from edge")
        # work on requests
        for request in request_iterator:
            if not request.ListFields():
                print("Data empty!")
                continue
            request_sensor_type = request.data.type
            request_sensor_value = request.data.value
            # create task, add task to queue
            task = Task(value=request_sensor_value, sensor_type=request_sensor_type)
            self.task_queue.add_task(task)
            print(f"Received data from edge and added task to queue ", task.id)

            # send finished tasks to edge
            done_tasks = self.task_queue.get_all_processed_tasks()
            for task in done_tasks:
                print(f"Sending feedback for task: {task}")
                yield self.send_feedback(task)

    def get_first_unprocessed_task(self):
        for task in self.task_queue.tasks:
            if not task.processed:
                return task
        return None

    def process_task_queue(self):
        while True:
            task = self.get_first_unprocessed_task()
            if task is not None:
                self.process_task(task=task)
                
            time.sleep(1)

    def process_task(self, task: Task):
        task_value = float(task.value)
        result_data = {}
        if task.sensor_type == 1:
            result_data = {'x':float(1 * task_value), 'y':float(1 * task_value/2), 'z':float(1 * task_value)}
        else:
            result_data = {'x':float(1 + task_value), 'y':float(1 + task_value/2), 'z':float(1 + task_value)}
        self.task_queue.update_task(task, "result", result_data)
        self.task_queue.update_task(task, "processed", True)
        
    def send_feedback(self, task):
        result_data = task.result
        pos =  fog_pb2.Position(x=result_data.get('x'), y=result_data.get('y'), z=result_data.get('z'))
        return fog_pb2.UpdatePositionResponse(position=pos)
            
def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cloud_service = CloudService()
    fog_pb2_grpc.add_CloudServiceServicer_to_server(cloud_service, server)

    SERVICE_NAMES = (
        fog_pb2.DESCRIPTOR.services_by_name["CloudService"].full_name,
        reflection.SERVICE_NAME,
    )

    # Enable reflection (for grpcui and grpcurl)
    reflection.enable_server_reflection(SERVICE_NAMES, server)

    server.add_insecure_port("[::]:50051")
    server.start()
    print("Server started on port 50051")

    task_processor_thread = threading.Thread(
        target=cloud_service.process_task_queue
    )
    task_processor_thread.start()
    print("Started Queue Processor on separate thread")

    server.wait_for_termination()


if __name__ == "__main__":
    serve()
