/**
* SkillChain - Chaincode methods for Candidates
* @author Akhil Mohandas (akhilmohandas@creopedia.com)
* Copyright creopedia business intelligence. 2019 All Rights Reserved.
**/


package main


import (
	"fmt"
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)



// createCandidate - Creates a Candidate Record
// @params {candidate_id, name, gender, date_of_birth, phone, mail,address,
// fathername,fatheroccupation,fatherphone,mothername,mother occupation,motherphone,idcard_number}
// @params address {address_line_1, street, city, state, country, postal_code}
func (t *skillChainChaincode) createCandidate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) !=3{
		logger.Infof("Incorrect number of arguments. Expecting 3 arguments to create candidate record. got %s\n",string(len(args)))
		return shim.Error("Incorrect number of arguments. Expecting 3 arguments to create candidate record.")

	}

	var candidateRecord Candidate
	var candidateID string

	candidateID = args[0]
	
	// Checking whether candidate is already created
	candidateRecordsCheck, err := stub.GetState(candidateID)

	if err != nil || candidateRecordsCheck != nil {
		return shim.Error("candidate id is already exist in the system")
	}

	json.Unmarshal([]byte(args[1]), &candidateRecord)

	candidateAsBytes, _ := json.Marshal(candidateRecord)

	
	stub.PutState(candidateID, candidateAsBytes)

	// Transaction Response
	logger.Infof("Create candidate Response:%s\n", string(candidateAsBytes))
	return shim.Success(nil)


}


//queryCandidate - query a candidate using the key
//@params key

func (t *skillChainChaincode) queryCandidate(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	var Key string
	var jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	Key = args[0]

	// Get the state from the ledger
	// patientRecords, err := stub.GetState(searchKey)
	candidateRecords, err := stub.GetState(Key)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + Key + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", Key))
	}

	if candidateRecords == nil {
		jsonResp = "{\"Error\":\"No data found for " + Key + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", Key))
	}

	jsonResp = "{\"Search Key\":\"" + Key + "\",\"Data\":\"" + string(candidateRecords) + "\"}"
	logger.Infof("Query Response: %s\n", jsonResp)
	return shim.Success(candidateRecords)


}



func (t *skillChainChaincode) updateCandidate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3 arguments to update candidate record, patient id, data and user email")
	}

	var candidate Candidate
	var UpdateCandidates UpdateCandidate
	var key string
	var orgnization string


	key = args[0]
	orgnization = args[2]

	if orgnization == "Org1" {
	// Get the state from the ledger
	candidatesBytes, err := stub.GetState(key)

	if err != nil || candidatesBytes == nil {
		return shim.Error(fmt.Sprintf("No data found for the key %s",key))
	}
	json.Unmarshal(candidatesBytes, &candidate)
	json.Unmarshal([]byte(args[1]), &UpdateCandidates)


	updateCandidateAsBytes, _ := json.Marshal(UpdateCandidates)
	json.Unmarshal(updateCandidateAsBytes, &candidate)
	candidateAsBytes,_ := json.Marshal(candidate)
	
	
	stub.PutState(key, candidateAsBytes)


	// Transaction Response
	logger.Infof("Update Candidate Response:%s\n", string(candidateAsBytes))
	return shim.Success(nil)
} else{
	logger.Infof("Not authorized to update the record")
	return shim.Error("Not authorized to update the record")
}
}
