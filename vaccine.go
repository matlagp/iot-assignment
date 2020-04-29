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


type QueryExperimentResult struct {
	Key    string `json:"Key"`
	Record *VaccineExperiment
}


type QueryConclusionResult struct {
	Key    string `json:"Key"`
	Record *Conclusion
}


func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	led1 := Person { Name: "John", Surname: "Smith", Id: 1}
	led2 := Person { Name: "Elizabeth", Surname: "Winter", Id: 2}
	led3 := Person { Name: "Jim", Surname: "Badley", Id: 3}

	res1 := []Person {Person { Name: "Fei", Surname: "Chu", Id: 4},
						Person { Name: "Helena", Surname: "Gardner", Id: 5},
						Person { Name: "Simon", Surname: "Sutton", Id: 6}}
	res2 := []Person {Person { Name: "Milo", Surname: "Pacher", Id: 7},
						Person { Name: "Andrew", Surname: "Human", Id: 8}}
	res3 := []Person {Person { Name: "Nicolas", Surname: "Contino", Id: 9}}

	t1s, _ = time.Parse(shortForm, "2020-Feb-03")
	t1e, _ = time.Parse(shortForm, "2020-May-20")
	t2s, _ = time.Parse(shortForm, "2018-Jan-15")
	t2e, _ = time.Parse(shortForm, "2019-Jan-01")
	t3s, _ = time.Parse(shortForm, "2017-Dec-05")
	t3e, _ = time.Parse(shortForm, "2018-Nov-03")

	exps := []VaccineExperiment{
		VaccineExperiment{Company: "International Medical Laboratory", Description: "Testing vaccine for coronavirus",
							Leader: led1, Researchers: res1, VaccineName: "VAC-COV-1", Disease: "COVID-19",
							Conclusions := []Conclusion, StartTime := t1s, EndTime := t1e },
						Company: "New York Medical Labolatory", Description: "Testing vaccine for new mutations of common flu.",
							Leader: led2, Researchers: res2, VaccineName: "FLU-42-B", Disease: "Common flu",
							Conclusions := []Conclusion, StartTime := t2s, EndTime := t2e },
						Company: "California Main Hospital", Description: "Experimenting with vaccine for well known disease, which is laziness",
							Leader: led3, Researchers: res3, VaccineName: "ANTI-LAZY-v0", Disease: "Laziness",
							Conclusions := []Conclusion, StartTime := t3s, EndTime := t3e }
	}

	for i, exp := range exps {
		expAsBytes, _ := json.Marshal(exp)
		err := ctx.GetStub().PutState("EXP"+strconv.Itoa(i), expAsBytes)

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
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
