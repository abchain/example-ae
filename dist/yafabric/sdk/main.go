package main

import (
	_ "hyperledger.abchain.org/adapter/yafabric/client"
	"hyperledger.abchain.org/example/ae/service"
)

func main() {

	// Start SDK Service
	service.StartService()
}
