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


// createEducation - Creates an Education Record
// @params {candidate_id, name, gender, date_of_birth, phone, mail,address,
//fathername,fatheroccupation,fatherphone,mothername,mother occupation,motherphone,idcard_number}
// @params address {address_line_1, street, city, state, country, postal_code}
func (t *skillChainChaincode) createEducation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	if len(args) !=2{
		logger.Infof("Incorrect number of arguments. Expecting 2 arguments to create Education record. got %s\n",string(len(args)))
		return shim.Error("Incorrect number of arguments. Expecting 2 arguments to create Education record.")
	}


	var educationRecord Education
	var educationID string

	educationID = string('E') + args[0]


	// Checking whether Education is already created
	educationRecordsCheck, err := stub.GetState(educationID)

	if err != nil || educationRecordsCheck != nil {
		return shim.Error("Education record for the candidate is already exist in the system")
	}

	json.Unmarshal([]byte(args[1]), &educationRecord)

	educationAsBytes, _ := json.Marshal(educationRecord)

	
	stub.PutState(educationID, educationAsBytes)

	// Transaction Response
	logger.Infof("Create candidate Response:%s\n", string(educationAsBytes))
	return shim.Success(nil)

}


//queryEducation - query a candidate using the key
//@params key

func (t *skillChainChaincode) queryEducation(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	var Key string
	var jsonResp string
	var err error
	var edu Education

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting a search key")
	}

	Key = args[0]

	// Get the state from the ledger
	educationRecord, err := stub.GetState(Key)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + Key + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", Key))
	}

	if educationRecord == nil {
		jsonResp = "{\"Error\":\"No data found for " + Key + "\"}"
		logger.Infof("Query Response:%s\n", jsonResp)
		return shim.Error(fmt.Sprintf("Failed to get state for the key %s", Key))
	}

	json.Unmarshal(educationRecord, &edu)

	jsonResp = "{\"Search Key\":\"" + Key + "\",\"Data\":\"" + string(educationRecord) + "\"}"
	logger.Infof("Query Response: %s\n", jsonResp)
	/*resArray := [2]string{edu.TenthBoard,edu.TenthMark}
	res,_ := json.Marshal(resArray)*/

	return shim.Success(educationRecord)


}
