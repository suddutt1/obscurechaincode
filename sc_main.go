package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var _MAIN_LOGGER = shim.NewLogger("SmartContractMain")

// Init initializes chaincode.
func (sc *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	_MAIN_LOGGER.Infof("Inside the init method ")
	response := sc.init(stub)
	return response
}

//Invoke is the entry point for any transaction
func (sc *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, _ := stub.GetFunctionAndParameters()
	switch function {
	case "callExternal":
		return sc.callExternal(stub)
	}
	return sc.handleFunctions(stub)
}
func (sc *SmartContract) callExternal(stub shim.ChaincodeStubInterface) pb.Response {
	postBody := make(map[string]interface{})
	postBody["field1"] = "DATA 1"
	postBody["field2"] = "DATA 2"
	postBody["field3"] = "DATA 3"

	isCallGood, respBytes := PostDataWithResponse("http://external.api.net/api/", postBody)
	if isCallGood {
		return shim.Success(respBytes)
	}
	return shim.Error("Call to the external API failed ")
}
func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		_MAIN_LOGGER.Criticalf("Error starting  chaincode: %v", err)
	}
}
