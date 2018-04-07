/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/Pallinder/go-randomdata"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Salmon struct {
	Vessel   string `json:"vessel"`
	Datetime string `json:"datetime"`
	Location string `json:"location"`
	Holder   string `json:"holder"`
}

type SalmonChaincode struct {
}

func (t *SalmonChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t *SalmonChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	var result []byte
	var err error
	switch fn {
	case "initLedger":
		result, err = initLedger(stub, args)
	case "recordSalmon":
		result, err = recordSalmon(stub, args)
	case "changeSalmonHolder":
		result, err = changeSalmonHolder(stub, args)
	case "querySalmon":
		result, err = querySalmon(stub, args)
	case "queryAllSalmon":
		result, err = queryAllSalmon(stub, args)
	default:
		return shim.Error("Unsupported function: " + fn)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success(result)
}

func initLedger(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	spawnCount := 100
	if len(args) == 1 {
		var err error
		spawnCount, err = strconv.Atoi(args[0])
		if err != nil {
			return nil, fmt.Errorf("read param spawnCount fail: %s", err)
		}
	}

	for i := 1; i <= spawnCount; i++ {
		_, err := recordSalmon(stub, []string{
			strconv.Itoa(i),
			randomdata.SillyName(),
			randomdata.FullDateInRange("2016-08-01", "2018-08-22"),
			randomdata.City(),
			"fredrick",
		})
		if err != nil {
			return nil, fmt.Errorf("Record salmon fail: %s", err)
		}
	}

	return nil, nil
}

func recordSalmon(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 5 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting 5")
	}

	id := args[0]
	vessel := args[1]
	datetime := args[2]
	location := args[3]
	holder := args[4]

	salmon := Salmon{Vessel: vessel, Datetime: datetime, Location: location, Holder: holder}
	data, err := json.Marshal(salmon)
	if err != nil {
		return nil, fmt.Errorf("marshal fail: %s", err)
	}

	err = stub.PutState(id, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to set asset: %s", args[0])
	}

	return nil, nil
}

func changeSalmonHolder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting 2")
	}

	id := args[0]
	holder := args[1]

	data, err := stub.GetState(id)
	if err != nil {
		return nil, fmt.Errorf("get salmon fail: %s", err)
	}

	var salmon Salmon
	err = json.Unmarshal(data, &salmon)
	if err != nil {
		return nil, fmt.Errorf("unmarshal fail: %s", err)
	}

	salmon.Holder = holder

	data, err = json.Marshal(salmon)
	if err != nil {
		return nil, fmt.Errorf("marshal fail: %s", err)
	}

	err = stub.PutState(id, data)
	if err != nil {
		return nil, fmt.Errorf("Failed to set asset: %s", args[0])
	}

	return nil, nil
}

func querySalmon(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting 1")
	}

	id := args[0]
	data, err := stub.GetState(id)
	if err != nil {
		return nil, fmt.Errorf("get salmon fail: %s", err)
	}

	return marshalSalmonDocument(id, data), nil
}

func marshalSalmonDocument(id string, data []byte) []byte {
	var salmon Salmon
	json.Unmarshal(data, &salmon)

	document := struct {
		ID     string `json:"id"`
		Salmon `json:",inline"`
	}{id, salmon}

	documentData, _ := json.Marshal(document)

	return documentData
}

func queryAllSalmon(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) > 2 {
		return nil, fmt.Errorf("Incorrect arguments. Expecting 0 to 2")
	}

	var startKey, endKey string
	if len(args) > 0 {
		startKey = args[0]
	}
	if len(args) > 1 {
		endKey = args[1]
	}

	iter, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return nil, fmt.Errorf("get salmons fail: %s", err)
	}
	defer iter.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for iter.HasNext() {
		queryResponse, err := iter.Next()
		if err != nil {
			return nil, fmt.Errorf("retrive next salmon fail: %s", err)
		}

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(marshalSalmonDocument(queryResponse.Key, queryResponse.Value)))
		bArrayMemberAlreadyWritten = true
	}

	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SalmonChaincode)); err != nil {
		fmt.Printf("Error starting SalmonChaincode: %s", err)
	}
}
