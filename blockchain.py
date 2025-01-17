#!/usr/bin/env python
# -*- coding: utf-8 -*-

'''
  bcosliteclientpy is a python client for FISCO BCOS2.0 (https://github.com/FISCO-BCOS/)
  bcosliteclientpy is free software: you can redistribute it and/or modify it under the
  terms of the MIT License as published by the Free Software Foundation. This project is
  distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even
  the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. Thanks for
  authors and contributors of eth-abi, eth-account, eth-hash，eth-keys, eth-typing, eth-utils,
  rlp, eth-rlp , hexbytes ... and relative projects
  @author: kentzhang
  @date: 2019-06
'''

from client.contractnote import ContractNote
from client.bcosclient import BcosClient
import os
from eth_utils import to_checksum_address
from client.datatype_parser import DatatypeParser
from client.common.compiler import Compiler
from client.bcoserror import BcosException, BcosError
from client_config import client_config

# 从文件加载abi定义
if os.path.isfile(client_config.solc_path) or os.path.isfile(client_config.solcjs_path):
    Compiler.compile_file("contracts/v1_2.sol")
abi_file = "contracts/v1_2.abi"
data_parser = DatatypeParser()
data_parser.load_abi_file(abi_file)
contract_abi = data_parser.contract_abi

try:
    client = BcosClient()
    print(client.getinfo())
    # 部署合约
    # print("\n>>Deploy:----------------------------------------------------------")
    # with open("contracts/v1_2.bin", 'r') as load_f:
    #     contract_bin = load_f.read()
    #     load_f.close()
    # result = client.deploy(contract_bin)
    # print("deploy", result)
    # print("new address : ", result["contractAddress"])
    contract_name = os.path.splitext(os.path.basename(abi_file))[0]
    # memo = "tx:" + result["transactionHash"]
    # 把部署结果存入文件备查
    # ContractNote.save_address(contract_name,
    #                           result["contractAddress"],
    #                           int(result["blockNumber"], 16), memo)
    # 发送交易，调用一个改写数据的接口
    print("\n>>sendRawTransaction:----------------------------------------------------")
    # 获取智能合约的地址
    to_address = ContractNote.get_last(contract_name)
    args = None

    receipt = client.sendRawTransactionGetReceipt(to_address, contract_abi, "Task1", args)
    print("receipt:", receipt)

    # 解析receipt里的log
    print("\n>>parse receipt and transaction:--------------------------------------")
    txhash = receipt['transactionHash']
    print("transaction hash: ", txhash)
    logresult = data_parser.parse_event_logs(receipt["logs"])
    i = 0
    for log in logresult:
        if 'eventname' in log:
            i = i + 1
            print("{}): log name: {} , data: {}".format(i, log['eventname'], log['eventdata']))
    # 获取对应的交易数据，解析出调用方法名和参数

    txresponse = client.getTransactionByHash(txhash)
    inputresult = data_parser.parse_transaction_input(txresponse['input'])
    print("transaction input parse:", txhash)
    print(inputresult)

    # 解析该交易在receipt里输出的output,即交易调用的方法的return值
    outputresult = data_parser.parse_receipt_output(inputresult['name'], receipt['output'])
    print("receipt output :", outputresult)
except BcosException as e:
    print("execute demo_transaction failed for: {}".format(e))
except BcosError as e:
    print("execute demo_transaction failed for: {}".format(e))