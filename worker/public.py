from client.contractnote import ContractNote
from client.bcosclient import BcosClient
import os
import json
from eth_utils import to_checksum_address
from client.datatype_parser import DatatypeParser
from client.common.compiler import Compiler
from client.bcoserror import BcosException, BcosError
from client_config import client_config
import gateway_pb2

abi_file = "contracts/v1_2.abi"

def EventEmitter(stub, initHeight):
    data_parser = DatatypeParser()
    try:
        client = BcosClient()
        lastHeight = initHeight
        while 1:
            currHeight = client.getBlockNumber()
            if currHeight > lastHeight:
                for height in range(lastHeight + 1, currHeight + 1):
                    result = client.getBlockByNumber(height, _includeTransactions=False)
                for hash in result["transactions"]:
                    receipt = client.getTransactionReceipt(hash)
                    logs = data_parser.parse_event_logs(receipt["logs"])
                    for log in logs:
                        print(log)
                        if log["eventname"] == "Task1Finished":
                            stub.PublishMessage(gateway_pb2.PublishMessageRequest(name="task1", correlationKey="task1", variables=json.dumps({"contract": receipt["to"]})))
                        elif log["eventname"] == "Task2Finished":
                            stub.PublishMessage(gateway_pb2.PublishMessageRequest(name="task2", correlationKey=receipt["to"]))
    except BcosException as e:
        print("execute demo_transaction failed for: {}".format(e))
    except BcosError as e:
        print("execute demo_transaction failed for: {}".format(e))

def deploySmartContract():
    data_parser = DatatypeParser()
    data_parser.load_abi_file(abi_file)
    contract_abi = data_parser.contract_abi
    try:
        client = BcosClient()
        # 部署合约
        with open("contracts/v1_2.bin", 'r') as load_f:
            contract_bin = load_f.read()
            load_f.close()
        result = client.deploy(contract_bin)
        # init
        client.sendRawTransactionGetReceipt(result["contractAddress"], contract_abi, "Init")
        return result["contractAddress"]
    except BcosException as e:
        print("execute demo_transaction failed for: {}".format(e))
    except BcosError as e:
        print("execute demo_transaction failed for: {}".format(e))