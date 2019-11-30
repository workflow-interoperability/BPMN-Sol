package main

import (
	v1 "bpmn-sol/v1"
	"fmt"
	"log"

	"github.com/FISCO-BCOS/go-sdk/accounts/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/common"
	"github.com/FISCO-BCOS/go-sdk/crypto"
)

func main() {
	groupID := uint(1)
	client, err := client.Dial("http://localhost:8545", groupID)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("input your privateKey in hex without \"0x\"") // 145e247e170ba3afd6ae97e88f00dbc976c2345d511b0f6713355d19d8b80b58
	if err != nil {
		log.Fatal(err)
	}
	// deploy
	auth := bind.NewKeyedTransactor(privateKey) // input your privateKey
	address, tx, instance, err := v1.DeployV1(auth, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract address: ", address.Hex()) // the address should be saved
	fmt.Println("transaction hash: ", tx.Hash().Hex())
	// load
	address = common.HexToAddress("contract addree in hex") // 0x0626918C51A1F36c7ad4354BB1197460A533a2B9
	instance, err = v1.NewV1(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("contract is loaded")
	// call
	// init
	_, err = instance.Init(auth)
	if err != nil {
		log.Fatal(err)
	}
	// call task1
	_, err = instance.Task1(auth)
	if err != nil {
		log.Fatal(err)
	}
	// call task2
	_, err = instance.Task2(auth)
	if err != nil {
		log.Fatal(err)
	}
}
