import gateway_pb2
from worker import public
from client.datatype_parser import DatatypeParser
from client.bcosclient import BcosClient
import json

def Task1(job, stub):
    print(job)
    data_parser = DatatypeParser()
    data_parser.load_abi_file(public.abi_file)
    contract_abi = data_parser.contract_abi
    client = BcosClient()

    address = public.deploySmartContract()
    receipt = client.sendRawTransactionGetReceipt(address, contract_abi, "Task1")
    stub.SetVariables(gateway_pb2.SetVariablesRequest(elementInstanceKey=job.workflowInstanceKey, variables=json.dumps({"contract": address, "tx": receipt["transactionHash"]})))
    if bool(receipt["output"]):
        stub.CompleteJob(gateway_pb2.CompleteJobRequest(jobKey = job.key))
    else:
        stub.FailJob(gateway_pb2.FailJobRequest(jobKey = job.key))
        