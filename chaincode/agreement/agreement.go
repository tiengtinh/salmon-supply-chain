/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Agreement struct {
	Price float64 `json:"price"`
}

type AgreementChaincode struct {
}

func (t *AgreementChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *AgreementChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result []byte
	var err error
	switch fn {
	case "recordAgreement":
		result, err = recordAgreement(stub, args)
	case "queryAgreement":
		result, err = queryAgreement(stub, args)
	default:
		return shim.Error("Unsupported function: " + fn)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success(result)
}

func recordAgreement(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting 2")
	}
	id := args[0]
	price, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return nil, fmt.Errorf("invalid price[%s]: %s", args[1], err)
	}

	agreement := Agreement{price}
	data, err := json.Marshal(agreement)
	if err != nil {
		return nil, fmt.Errorf("marshal fail: %s", err)
	}

	err = stub.PutState(id, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to put state: %s", args[0])
	}

	return nil, nil
}

func queryAgreement(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting 1")
	}
	id := args[0]

	data, err := stub.GetState(id)
	if err != nil {
		return nil, fmt.Errorf("get agreement fail: %s", err)
	}

	return data, nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(AgreementChaincode)); err != nil {
		fmt.Printf("Error starting AgreementChaincode: %s", err)
	}
}
