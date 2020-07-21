package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	hycc "hyperledger.abchain.org/adapter/hyfabric/chaincode"
	"hyperledger.abchain.org/example/ae/chaincode/cc"
)

func main() {
	err := shim.Start(hycc.AdaptChainCode(chaincode.NewChaincode(true)))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
