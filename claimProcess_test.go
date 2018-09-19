/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

// ====CHAINCODE EXECUTION SAMPLES (BCS REST API) ==================
/*
#TEST transaction / Init ledger

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerB","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerC","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initLedgerD","args":["ser1234"],"chaincodeVer":"v1"}'

 # TEST transaction / Add Car Part

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1234", "tata", "1502688979", "airbag 2020", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1235", "tata", "1502688979", "airbag 2020", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1236", "tata", "1502688979", "airbag 2020", "toyota", "false", "15026889790"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1237", "tata", "1502688979", "airbag 5000", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1238", "tata", "1502688979", "airbag 5000", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehiclePart","args":["ser1239", "tata", "1502688979", "airbag 5000", "toyota", "false", "15026889790"],"chaincodeVer":"v1"}'

# TEST transaction / Add Car

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["maz1000001", "mazda", "mazda 6", "1502688979", "ser1235", "mazda", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["ren1000001", "renault", "megan", "1502688979", "ser1236", "renault", "false", "1502688979"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"initVehicle","args":["ford1000001", "ford", "mustang", "1502688979", "ser1237", "ford", "false", "1502688979"],"chaincodeVer":"v1"}'

# TEST query / Populated database

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"readVehiclePart","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"readVehicle","args":["mer1000001"],"chaincodeVer":"v1"}'

# TEST transaction / Transfer ownership

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"transferVehiclePart","args":["ser1234", "mercedes"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"transferVehicle","args":["mer1000001", "mercedes los angeles"],"chaincodeVer":"v1"}'

# TEST query / Get History

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"getHistoryForRecord","args":["ser1234"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/query -d '{"channel":"channel1","chaincode":"vehiclenet","method":"getHistoryForRecord","args":["mer1000001"],"chaincodeVer":"v1"}'

# TEST transaction / delete records

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"deleteVehiclePart","args":["ser1235"],"chaincodeVer":"v1"}'
curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/rest/v1/transaction/invocation -d '{"channel":"channel1","chaincode":"vehiclenet","method":"deleteVehicle","args":["maz1000001"],"chaincodeVer":"v1"}'

# TEST transaction / Recall Part

curl -H "Content-type:application/json" -X POST http://localhost:3100/bcsgw/{"channel":"channel1","chaincode":"vehiclenet","method":"setPartRecallState","args":["abg1234",true],"chaincodeVer":"v3"}'

 */

package main

import (
//	"bytes"
//	"encoding/json"
	"fmt"
/*	"strconv"
	"strings"
	"time"
	*/
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
//	"github.com/hyperledger/fabric/protos/peer"

)

//ClaimProcessChaincode example simple Chaincode implementation
type ClaimProcessChaincodeTest struct {
}

// @claimPart JSON object
type claimInfoTest struct {
	ObjectType   				string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	SerialNumber 				string `json:"serialNumber"`
	InsuredName  				string `json:"insuredName"`
	PolicyID     				string `json:"policyID"`
	ClaimID      				string `json:"claimID"`
	InsuranceStartDate 	int    `json:"insuranceStartDate"`
	InsuranceEndDate 		int    `json:"insuranceEndDate"`
	AdmissionDate 			int    `json:"admissionDate"`
	DischargeDate 			int    `json:"dischargeDate"`
	PatientName  				string `json:"patientName"`
	PatientGender  			string `json:"patientGender"`
	DOB 								int 	 `json:"dob"`
	Relationship  			string `json:"relationship"`
	ClaimType  					string `json:"claimType"`
	HospitalName  			string `json:"hospitalName"`
	DoctorName  				string `json:"doctorName"`
	NatureOfIllness 		string `json:"natureOfIllness"`
	Diagnosis 					string `json:"diagnosis"`
	Accident       			bool   `json:"accident"`
	FIR       					bool   `json:"fir"`
	FIRNumber						string `json:"firNumber"`
	PoliceStation  			string `json:"policeStation"`
	OtherPolicy       	bool   `json:"otherPolicy"`
	OtherPolicyId       string `json:"otherPolicyId"`
	ClaimStatus  				string `json:"claimStatus"`
}

// @the bill information type is below
type billInfoTest struct {
	ObjectType         string 	`json:"docType"`       //docType is used to distinguish the various types of objects in state database
	ClaimID      			 string 	`json:"claimID"`
	ClaimItemId      	 string 	`json:"claimItemId"`
	Description        string 	`json:"description"`
	BillNo             string 	`json:"billNo"`
	BillDate       		 int   	  `json:"billDate"`
	BillType 					 string 	`json:"billType"`
	Amount             string 	`json:"amount"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(ClaimProcessChaincode))
	if err != nil {
		fmt.Printf("Error starting claim process chaincode: %s", err)
	}
}

/*
// Init initializes chaincode
// ===========================
func (t *ClaimProcessChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
} */

func TestInitPartCreatesIndexes(t *testing.T) {
	stub := shim.NewMockStub("mockChaincodeStub", new(ClaimProcessChaincode))
	if stub == nil {
			t.Fatalf("MockStub creation failed")
		}
		 args := [][]byte{
			    []byte("initclaimProcess"), []byte("sr1234"),
					[]byte("Deepak"), []byte("Pol1"),
					[]byte("C1"), []byte("1502688979"),
					[]byte("1502688979"), []byte("1502688979"),
					[]byte("1502688979"), []byte("abc"),
					[]byte("Female"), []byte("1502688979"),
					[]byte("Wife"), []byte("Hospitalization"),
					[]byte("City"), []byte("Dr Alok"),
					[]byte("Fever"), []byte("Dengue"),
					[]byte("false"), []byte("false"),
					[]byte("0"), []byte("0"),
					[]byte("false"), []byte("0"),
					[]byte("Applied")}
					invokeResult := stub.MockInvoke("12345", args)
					fmt.Println("-invoking result ")
					if invokeResult.Status != 200 {
						t.Errorf("Create Part returned non-OK status, got: %d, want: %d.", invokeResult.Status, 200)
					}
					//Validate index was created
					indexKey, _ := stub.CreateCompositeKey("claimInfo~claimID", []string{"Pol1","C1"})
					fmt.Println("-invoking index entry ")
					indexEntry, _ := stub.GetState(indexKey)
		 			if indexEntry == nil {
							t.Errorf("Index on assembler not created!")
							}
	}

/*
// Invoke - Our entry point for Invocations
// ========================================
func (t *ClaimProcessChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initclaimProcess" { //create a new Claim
		return t.initclaimProcess(stub, args)
	}

	/*else if function == "addDiagnosisInfo" { //update illness information
		return t.addDiagnosisInfo(stub, args)
	} else if function == "updateBillInfo" { //update bill information
		return t.updateBillInfo(stub, args)
	}else if function == "deleteClaim" { //delete a Claim
		return t.deleteClaim(stub, args)
	} else if function == "readClaim" { //read a Claim
		return t.readClaim(stub, args)
	} else if function == "getHistoryForClaim" { //get history of values for a claim
		return t.getHistoryForPolicy(stub, args)
	} else if function == "gethistoryForPolicy" { //get history of values for a Policy
		return t.getBillsByClaim(stub, args)
	} else if function == "getClaimStatus" { //get claim Status
		return t.getClaimStatus(stub, args)
	} else if function == "updateClaimStatus" { //update Claim Status
		return t.updateClaimStatus(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initVehiclePart - create a new vehicle part, store into chaincode state
// ============================================================
func (t *ClaimProcessChaincode) initclaimProcess(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	// sample data model with
	//   0       	1      		2     		3				4						5	  6
	// "ser1234", "Raj", "grpHealth1", "1", "1502688979", "1502688999", "1502688989","1502688990","Sharmila", "Female", "1502688969", "Wife", "Hospitalization", "Apollo", "Dr. Arvind", "Fever", "Dengue" , "No", "","","","", "","","","", "admission"
	if len(args) < 4  {
		return shim.Error("Incorrect number of arguments. Expecting atleast 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init claim process")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7th argument must be a non-empty string")
	}
	if len(args[7]) <= 0 {
		return shim.Error("8th argument must be a non-empty string")
	}
	if len(args[8]) <= 0 {
		return shim.Error("9th argument must be a non-empty string")
	}
	if len(args[9]) <= 0 {
		return shim.Error("10th argument must be a non-empty string")
	}
	if len(args[10]) <= 0 {
		return shim.Error("11th argument must be a non-empty string")
	}
	serialNumber := args[0]
	insuredName := strings.ToLower(args[1])
	policyID := strings.ToLower(args[2])
	claimID := strings.ToLower(args[3])
	insuranceStartDate, err := strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("4rd argument must be a numeric string")
	}
	insuranceEndDate, err := strconv.Atoi(args[5])
	if err != nil {
		return shim.Error("5th argument must be a numeric string")
	}
	admissionDate, err := strconv.Atoi(args[6])
	if err != nil {
		return shim.Error("6th argument must be a numeric string")
	}
	dischargeDate, err := strconv.Atoi(args[7])
	if err != nil {
		return shim.Error("7th argument must be a numeric string")
	}
	patientName := strings.ToLower(args[8])
	patientGender := strings.ToLower(args[9])
	dob, err := strconv.Atoi(args[10])
	if err != nil {
		return shim.Error("11th argument must be a numeric string")
	}

	relationship:= strings.ToLower(args[11])
	claimType:= strings.ToLower(args[12])
	hospitalName:= strings.ToLower(args[13])
	doctorName:= strings.ToLower(args[14])
	natureOfIllness:= strings.ToLower(args[15])
	diagnosis:= strings.ToLower(args[16])
	accident, err := strconv.ParseBool(args[17])
	if err != nil {
		return shim.Error("16th argument must be a boolean string")
	}

	fir, err := strconv.ParseBool(args[18])
	if err != nil {
		return shim.Error("17th argument must be a boolean string")
	}

	firNumber:= strings.ToLower(args[19])
	policeStation:= strings.ToLower(args[20])
	otherPolicy, err := strconv.ParseBool(args[21])

	if err != nil {
		return shim.Error("20th argument must be a boolean string")
	}
	otherPolicyId:= args[22]
	claimStatus:= strings.ToLower(args[23])

	// ==== Check if claim already exists ====
	claimAsBytes, err := stub.GetState(claimID)
	if err != nil {
		return shim.Error("Failed to get claim " + err.Error())
	} else if claimAsBytes != nil {
		fmt.Println("This claim already exists: " + claimID)
		return shim.Error("This claim already exists: " + claimID)
	}

	// ==== Create claimInfo object and marshal to JSON ====
	objectType := "claimInfo"
	claimInfo := &claimInfo{objectType, serialNumber,insuredName,policyID,claimID,insuranceStartDate,insuranceEndDate,admissionDate,dischargeDate,patientName,patientGender,dob,relationship,claimType,hospitalName,doctorName,natureOfIllness,diagnosis,accident,fir,firNumber,policeStation,otherPolicy,otherPolicyId,claimStatus }
	claimJSONasBytes, err := json.Marshal(claimInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save claim to state ===
	err = stub.PutState(claimID, claimJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== claim saved Return success ====
	fmt.Println("- end init claim part")
	return shim.Success(nil)
}


// ============================================================
// initVehicle - create a new bill Info , store into chaincode state
// ============================================================
func (t *ClaimProcessChaincode) initVehicle(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	// data model with recall fields
	//   0       		1      		2     		3			   4		5	       6			7
	// "mer1000001", "mercedes", "c class", "1502688979", "ser1234", "mercedes", "false", "1502688979"

	// @MODIFY_HERE extend to expect 8 arguements, up from 6
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init bill ")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}
	if len(args[5]) <= 0 {
		return shim.Error("6th argument must be a non-empty string")
	}
	if len(args[6]) <= 0 {
		return shim.Error("7h argument must be a non-empty string")
	}
	claimID := args[0]
	claimItemId:=args[1]
	description:=args[2]
	billNo:=args[3]
	billDate:= strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("5th argument must be a numeric string")
	}
	billType:=args[5]
	amount:=args[6]

	// ==== Check if claim already exists ====
	billInfoAsBytes, err := stub.GetState(billID)
	if err != nil {
		return shim.Error("Failed to get claim: " + err.Error())
	} else if claimAsBytes != nil {
		return shim.Error("This claim already exists: " + claimID)
	}

	// ==== Create claim object and marshal to JSON ====
	objectType := "billInfo"
	billInfo := &billInfo{objectType,claimID,claimItemId,description,billNo,billDate,billType,amount}
	vehicleJSONasBytes, err := json.Marshal(vehicle)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save bill to state ===
	err = stub.PutState(claimBill, billInfoJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// ==== billinfo saved and indexed. Return success ====
	fmt.Println("- end init bill Info")
	return shim.Success(nil)
}

// ===============================================
// createIndex - create search index for ledger
// ===============================================
func (t *ClaimProcessChaincode) createIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start create index")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all claims made by policy ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Save index entry to state. Only the key name is needed, no need to store a duplicate copy of object.
	//  Note - passing a 'nil' value will effectively delete the key from state, therefore we pass null character as value
	value := []byte{0x00}
	stub.PutState(indexKey, value)

	fmt.Println("- end create index")
	return nil
}


// ===============================================
// deleteIndex - remove search index for ledger
// ===============================================
func (t *ClaimProcessChaincode) deleteIndex(stub shim.ChaincodeStubInterface, indexName string, attributes []string) error {
	fmt.Println("- start delete index")
	var err error
	//  ==== Index the object to enable range queries, e.g. return all parts made by supplier b ====
	//  An 'index' is a normal key/value entry in state.
	//  The key is a composite key, with the elements that you want to range query on listed first.
	//  This will enable very efficient state range queries based on composite keys matching indexName~color~*
	indexKey, err := stub.CreateCompositeKey(indexName, attributes)
	if err != nil {
		return err
	}
	//  Delete index by key
	stub.DelState(indexKey)

	fmt.Println("- end delete index")
	return nil
}

// ===============================================
// readClaim - read a Claim from chaincode state
// ===============================================
func (t *ClaimProcessChaincode) readClaim(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var claimID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Claim ID of the Claim to query")
	}

	claimID = args[0]
	valAsbytes, err := stub.GetState(claimID) //get the claim from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + claimID + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Claim does not exist: " + claimID + "\"}"
		fmt.Println(jsonResp)
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ===============================================
// BillsByClaim - read a vehicle from chaincode state
// ===============================================
func (t *ClaimProcessChaincode) readClaim(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var claimID, billNo, jsonResp string
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting chassis number of the vehicle to query")
	}

	claimID = args[0]
	billNo = args[1]
	valAsbytes, err := stub.GetState(claimID,billNo) //get the bill from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + billNo + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Bill does not exist: " + billNo + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(valAsbytes)
}

// ==================================================
// deleteClaim - remove a Claim key/value pair from state
// ==================================================
func (t *ClaimProcessChaincode) deleteClaim(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var jsonResp string
	var claimInfo claimInfo
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	claimID := args[0]

	// to maintain the claimID index, we need to read the claim
	valAsbytes, err := stub.GetState(claimID) //get the claim from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + claimID + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Claim does not exist: " + claimID + "\"}"
		return shim.Error(jsonResp)
	}

	err = json.Unmarshal([]byte(valAsbytes), &ClaimInfoJSON)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to decode JSON of: " + serialNumber + "\"}"
		return shim.Error(jsonResp)
	}

	err = stub.DelState(claimID) //remove the claim from chaincode state
	if err != nil {
		return shim.Error("Failed to delete state:" + err.Error())
	}

	return shim.Success(nil)
}

// ===========================================================
// transfer a vehicle part by setting a new owner name on the vehiclePart
// ===========================================================
func (t *ClaimProcessChaincode) addDiagnosisInfo(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	//   0       1       3
	// "name", "from", "to"
	if len(args) < 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	claimID := args[0]
	diagnosis := strings.ToLower(args[1])
	natureOfIllness := strings.ToLower(args[2])
	fmt.Println("- start update Diagnosis Info ", claimID, diagnosis, natureOfIllness)

	message, err := t.addDiagnosisHelper(stub, claimID, diagnosis, natureOfIllness)
	if err != nil {
		return shim.Error(message + err.Error())
	} else if message != "" {
		return shim.Error(message)
	}

	fmt.Println("- end addDiagnosis (success)")
	return shim.Success(nil)
}

// ===========================================================
// addDiagnosis : helper method for addDiagnosis
// ===========================================================
/func (t *ClaimProcessChaincode) addDiagnosisHelper(stub shim.ChaincodeStubInterface, ClaimID string, diagnosis string, natureOfIllness string) (string, error) {
	// attempt to get the current claim object by serial number.
	// if sucessful, returns us a byte array we can then us JSON.parse to unmarshal
	fmt.Println("Transfering claim with Claim ID: " + claimID )
	claimInfoAsBytes, err := stub.GetState(claimID)
	if err != nil {
		return "Failed to get claim: " + claimID, err
	} else if vehiclePartAsBytes == nil {
		return "Claim does not exist: " + claimID, nil
	}

	claimInfoToTransfer := claimInfo{}
	err = json.Unmarshal(claimInfoAsBytes, &claimInfoToTransfer := claimInfo{} )
	//unmarshal it aka JSON.parse()
	if err != nil {
		return "", err
	}

	claimInfoToTransfer.diagnosis = diagnosis
  claimInfoToTransfer.natureOfIllness = natureOfIllness

	claimInfoJSONBytes, _ := json.Marshal(claimInfoToTransfer)
	err = stub.PutState(claimID, claimInfoJSONBytes) //rewrite the claimInfo
	if err != nil {
		return "", err
	}

	return "", nil
}

// ===========================================================================================
// getClaimByRange performs a range query based on the start and end keys provided.

// Read-only function results are not typically submitted to ordering. If the read-only
// results are submitted to ordering, or if the query is used in an update transaction
// and submitted to ordering, then the committing peers will re-execute to guarantee that
// result sets are stable between endorsement time and commit time. The transaction is
// invalidated by the committing peers if the result set has changed between endorsement
// time and commit time.
// Therefore, range queries are a safe option for performing update transactions based on query results.
// ===========================================================================================
func (t *ClaimProcessChaincode) getClaimByRange(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	startKey := args[0]
	endKey := args[1]

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getClaimByRange queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ===========================================================================================
// getHistoryForRecord returns the historical state transitions for a given key of a record
// ===========================================================================================
func (t *ProcessClaimChaincode) getHistoryForRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordKey := args[0]

	fmt.Printf("- start getHistoryForRecord: %s\n", recordKey)

	resultsIterator, err := stub.GetHistoryForKey(recordKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the key/value pair
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON vehiclePart)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getHistoryForRecord returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}


// ============================================================
// setPartRecallState - sets recall field of a vehicle
// ============================================================
func (t *AutoTraceChaincode) setPartRecallState(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var err error

	// expects following arguements
	//   	0       		1
	// "serialNumber", "status (boolean)"
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}

	serialNumber := args[0]
	recall, err := strconv.ParseBool(args[1])
	if err != nil {
		return shim.Error("2nd argument must be a boolean string")
	}

	// ==== Check if vehicle part already exists ====
	vehiclePartAsBytes, err := stub.GetState(serialNumber)
	if err != nil {
		return shim.Error("Failed to get vehicle part: " + err.Error())
	} else if vehiclePartAsBytes == nil {
		fmt.Println("This vehicle part does not exist: " + serialNumber)
		return shim.Error("This vehicle part does not exist:: " + serialNumber)
	}

	vehiclePartJSON := vehiclePart{}
	err = json.Unmarshal(vehiclePartAsBytes, &vehiclePartJSON) //unmarshal it aka JSON.parse()
	if err != nil {
		fmt.Println("Unable to unmarshall vehicle part from byte to JSON object: " + serialNumber)
		return shim.Error("Unable to unmarshall vehicle part from byte to JSON object: " + serialNumber)
	}

	// ==== Create vehiclePart object and marshal to JSON ====
	objectType := "vehiclePart"
	// Set now as the time of recall
	var recallTime int
	if recall {
		recallTime = int(time.Now().UTC().Unix())
	}else{
		recallTime = 0
	}

	vehiclePart := &vehiclePart{objectType, serialNumber, vehiclePartJSON.Assembler, vehiclePartJSON.AssemblyDate, vehiclePartJSON.Name, vehiclePartJSON.Owner, recall, recallTime}
	vehiclePartJSONasBytes, err := json.Marshal(vehiclePart)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save vehiclePart to state ===
	err = stub.PutState(serialNumber, vehiclePartJSONasBytes)

	// ==== Vehicle part saved. Return success ====
	fmt.Println("- end setPartRecallState")
	return shim.Success(nil)
}


*/
