package main

import (
	"hyperledger.abchain.org/cases/ae/service"
	_ "hyperledger.abchain.org/adapter/yafabric"
)

func main() {

	// Start SDK Service
	service.StartService()
}
