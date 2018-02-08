package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInit(t *testing.T, stub *shim.MockStub) {
	args := make([][]byte, 0)
	args = append(args, []byte("init"))
	res := stub.MockInit("1", args)
	if res.Status != shim.OK {
		fmt.Println("Init failed", string(res.Message))
		t.FailNow()
	}
}
func checkProbe(t *testing.T, stub *shim.MockStub) {
	args := make([][]byte, 0)
	args = append(args, []byte("probe"))

	res := stub.MockInvoke("1", args)
	if res.Status != shim.OK {
		t.FailNow()
	} else {
		fmt.Printf("\n %s\n", string(res.Payload))
	}
}

//Test_Init tests the input
func Test_Init(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("hcdm", scc)

	// Init A=123 B=234
	checkInit(t, stub)

}
func Test_Probe(t *testing.T) {
	scc := new(SmartContract)
	stub := shim.NewMockStub("hcdm", scc)
	checkInit(t, stub)
	checkProbe(t, stub)
}
