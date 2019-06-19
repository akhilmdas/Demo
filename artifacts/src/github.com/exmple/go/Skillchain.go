/*
* SkillChain - Chaincode
* @author Akhil Mohandas (akhilmohandas@creopedia.com)
* Copyright creopedia business intelligence. 2019 All Rights Reserved.
**/


package main


import (
	
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("SkillChain")

//skillChainChaincode chaincode implementation

type skillChainChaincode struct{
}


// Address data struct for address model
type Address struct {
	AddressLine1 string `json:"addressline1,omitempty"`
	Street       string `json:"street,omitempty"`
	City         string `json:"city,omitempty"`
	State        string `json:"state,omitempty"`
	Country      string `json:"country,omitempty"`
	PostalCode   string `json:"postalcode,omitempty"`
}


//Candidate data struct for storing candidate records
type Candidate struct {
	Name string `json:"name,omitempty"`
	Gender string `json:"gender,omitempty"`
	Address Address `json:"address,omitempty"`
	
}


//Education data struct for storing candidate records
type Education struct{
	TenthBoard string `json:"tenthBoard,omitempty"`
	TenthMark string `json:"tenthMark,omitempty"`
	TenthCertificate string `json:"tenthCertificate"`
}


//ExtraEducation data struct for storing candidate records
type ExtraEducation struct{
	Type string `json:"type,omitempty"`
	Acheivemnent string `json:"acheivement,omitempty"`
	Certificate string `json:"certificate,omitempty"`
}



// Function to initialize SmartContract in the network
func (t *skillChainChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### skillChain Init ###########")

	

	fmt.Println("Skillchain Chaincode Initilization Completed")
	return shim.Success(nil)
}


func (t *skillChainChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	function, args := stub.GetFunctionAndParameters()

	logger.Info("########### skillChain Invoke - " + function + " ###########")

	if function == "createCandidate" {
		return t.createCandidate(stub, args)
		// createCandidate - to Create Candidate Record
		// Creates Candidate Record with CandidateID as key.
	} else if function == "queryCandidate" {
		return t.queryCandidate(stub, args)
		// queryCandidate - to Query Candidate Record
		// Query Candidate Record with CandidateID as key.
	} else if function =="createEducation" {
		return t.createEducation(stub, args)
		// createEducation - to create Education record
		// creates Education record with educationId as key
	} else if function == "queryEducation" {
		return t.queryEducation(stub, args)
		// queryEducation - to query education record
		// query candidate record with educationid as key
	}
	

	logger.Errorf("unknown function name.  got: %v", args[0])
	return shim.Error(fmt.Sprintf("unknown function name. got: %v", args[0]))
}


// main function to start smart contract
func main() {
	err := shim.Start(new(skillChainChaincode))
	if err != nil {
		logger.Errorf("Error starting chaincode for skillchain: %s", err)
	}
}
