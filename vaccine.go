package main

import (
	"fmt"
	"time"
	"encoding/json"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

type Person struct {
	Name    string `json:”name”`
	Surname string `json:”surname”`
	Id      int    `json:”id”`
}

type Conclusion struct {
	Author       Person    `json:”author”`
	Content      string    `json:”content”`
	CreationTime time.Time `json:”time”`
}

type VaccineExperiment struct {
	Company     string       `json:"company"`
	Description string       `json:"description"`
	Leader      Person       `json:”leader”`
	Researchers []Person     `json:"researchers"`
	VaccineName string       `json:"vaccineName"`
	Disease     string       `json:"diseaseName"`
	Conclusions []Conclusion `json:”conclusions”`
	StartTime   time.Time    `json:”startTime”`
	EndTime     time.Time    `json:”endTime”`
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	return nil
}

func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error while creating chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error while starting chaincode: %s", err.Error())
	}
}
