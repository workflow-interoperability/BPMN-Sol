import grpc
import gateway_pb2
from grpc_gateway import gateway_pb2_grpc
import _thread
from worker import a
from worker import b
from worker import public
from client.bcosclient import BcosClient

channel = grpc.insecure_channel("localhost:26500")
stub = gateway_pb2_grpc.GatewayStub(channel)

def handleTask(stub, type, worker):
    for jobResponse in stub.ActivateJobs(gateway_pb2.ActivateJobsRequest(
        type = type, worker = worker, timeout=1000, maxJobsToActivate=32)):
        for job in jobResponse.jobs:
            if type == "task1":
                a.Task1(job, stub)
            else:
                b.Task2(job, stub)

client = BcosClient()
initHeight = client.getBlockNumber()

# handleTask(stub, "task1", "task1")
_thread.start_new_thread(handleTask, (stub, "task1", "task1"))
_thread.start_new_thread(handleTask, (stub, "task2", "task2"))
_thread.start_new_thread(public.EventEmitter, (stub, initHeight))

while 1:
   pass