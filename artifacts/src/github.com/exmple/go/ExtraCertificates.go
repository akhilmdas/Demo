/**
* SkillChain - Chaincode methods for Candidates
* @author Akhil Mohandas (akhilmohandas@creopedia.com)
* Copyright creopedia business intelligence. 2019 All Rights Reserved.
**/


package main

import(
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


// createExtraCertificates - create the record for extra certifications of the candidate
func (t *skillChainChaincode) createExtraCertificates(stub shim.ChaincodeStubInterface, args []string)pb.Response  {
	
	if len(args) != 2{
		logger.Infof("Incorrect number of arguments expecting 2")
		return shim.Error("Incorrect number of arguments expecting 2")
	}
	
	var ExtraEducation Extra
	var ExtraID string

	ExtraID = args[0]

	//checking whether the id already exist
	ExtraReccordChecks, err := stub.GetState(ExtraID)

	if err != nil || ExtraReccordChecks != nil {
		return shim.Error("Extra Education id already exists")
	}

	json.Unmarshal([]byte(args[1]), &Extra)
	extraAsBytes,_ := json.Marshal(Extra)

	//create the extra asset record

	stub.PutState(ExtraID, extraAsBytes)

	//Transaction Response
	logger.Infof("Create Record response :%s\n", string(extraAsBytes))
	return shim.Success(nil)
}

func (t *skillChainChaincode) queryExtra(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	var Key string
	var jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	Key = args[0]

	// Get the state from the ledger
	// patientRecords, err := stub.GetState(searchKey)
	extraRecords, err := stub.GetState(Key)

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

