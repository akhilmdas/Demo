package main
/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)




//Mychaincode define the smart contract structure
type Mychaincode struct{

}




//Candidate define the candiate structure
type Candidate struct{
	CandID string `json:"CandID"`
	Name string `json:"Name"`
	Phone string `json:"Phone"`
}


type Certificates struct{
	CertCandID string `json:"CertCandID"`
	certificate string `json:"Certificate"`
}




//Init is used to initialize the contract
func (s *Mychaincode) Init(stub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}




//Invoke is used to call different functions from the chaincode
func (s *Mychaincode) Invoke(stub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "setDetails" {
		return s.setDetails(stub, args)
	}else if function == "putDetails" {
		return s.putDetails(stub, args)
	}else if function == "setCertificate"{
		return s.setCertificate(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}




//SetDetials is used to save data to the blockchain
func (s *Mychaincode) setDetails(stub shim.ChaincodeStubInterface, args []string) sc.Response {

    if len(args) < 3 {
        return shim.Error("insert Into Table failed. Must include 3 column values")
    }

	var cand = Candidate{Name:args[1], Phone:args[2]}

	candAsBytes, _ := json.Marshal(cand)
	stub.PutState(args[0], candAsBytes)

    return shim.Success(nil)
}



func (s *Mychaincode) setCertificate(stub shim.ChaincodeStubInterface, args []string) sc.Response{

	if len(args)< 2{
		return shim.Error("invalid number of arguments, expected 2")
	}

	var cert = Certificates{certificate:args[1]}

	certAsBytes, _ := json.Marshal(cert)
	stub.PutState(args[0], certAsBytes)

	return shim.Success(nil)
}




//PutDetails is used for displaying the contentns fron the blockchain
func (s *Mychaincode) putDetails(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	candAsBytes, _ := stub.GetState(args[0])
	return shim.Success(candAsBytes)
}

func main() {
	var logger = shim.NewLogger("myChaincode")
	logger.SetLevel(shim.LogInfo)
	err := shim.Start(new(Mychaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}