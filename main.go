// Copyright Key Inside Co., Ltd. 2018 All Rights Reserved.

package main

import (
	"encoding/json"
	"math/big"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("knt-pci")

// Chaincode _
type Chaincode struct {
}

// Init implements shim.Chaincode interface.
func (cc *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke implements shim.Chaincode interface.
func (cc *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, params := stub.GetFunctionAndParameters()
	if txFn := routes[fn]; txFn != nil {
		return txFn(stub, params)
	}
	return shim.Error("unknown function: [" + fn + "]")
}

// TxFunc _
type TxFunc func(shim.ChaincodeStubInterface, []string) peer.Response

// routes is the map of invoke functions
var routes = map[string]TxFunc{
	"token": txToken,
	"mint":  txMint,
	"burn":  txBurn,
}

func txToken(stub shim.ChaincodeStubInterface, params []string) peer.Response {
	payload, err := json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(payload)
}

// params[0] : total supply
// params[1] : genesis account balance
// params[2] : max amount to mint (if val <= 0, no limit)
func txMint(stub shim.ChaincodeStubInterface, params []string) peer.Response {
	if len(params) != 3 {
		return shim.Error("incorrect number of parameters. expecting 3")
	}

	totalSupply, _ := big.NewInt(0).SetString(params[0], 10)
	balance, _ := big.NewInt(0).SetString(params[1], 10)
	amount, _ := big.NewInt(0).SetString(params[2], 10)
	available, err := mintableAmount(totalSupply, balance, amount)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(available.String()))
}

// params[0] : total supply
// params[1] : genesis account balance
// params[2] : max amount to burn (if val <= 0, no limit)
func txBurn(stub shim.ChaincodeStubInterface, params []string) peer.Response {
	if len(params) != 3 {
		return shim.Error("incorrect number of parameters. expecting 3")
	}

	totalSupply, _ := big.NewInt(0).SetString(params[0], 10)
	balance, _ := big.NewInt(0).SetString(params[1], 10)
	amount, _ := big.NewInt(0).SetString(params[2], 10)
	available, err := burnableAmount(totalSupply, balance, amount)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte(available.String()))
}

func main() {
	if err := shim.Start(new(Chaincode)); err != nil {
		logger.Criticalf("failed to start chaincode|%s", err)
	}
}
