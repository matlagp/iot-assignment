package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const shortForm = "2006-Jan-02"

type SmartContract struct {
	contractapi.Contract
}

type Person struct {
	Name    string `json:”name”`
	Surname string `json:”surname”`
	Licence string `json:”id”`
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
	Key     string `json:"Key"`
	Records []Conclusion
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	led1 := Person{Name: "John", Surname: "Smith", Licence: "R-001"}
	led2 := Person{Name: "Elizabeth", Surname: "Winter", Licence: "R-002"}
	led3 := Person{Name: "Jim", Surname: "Badley", Licence: "R-003"}

	res1 := []Person{Person{Name: "Fei", Surname: "Chu", Licence: "R-004"},
		Person{Name: "Helena", Surname: "Gardner", Licence: "R-005"},
		Person{Name: "Simon", Surname: "Sutton", Licence: "R-006"}}
	res2 := []Person{Person{Name: "Milo", Surname: "Pacher", Licence: "R-007"},
		Person{Name: "Andrew", Surname: "Human", Licence: "R-008"}}
	res3 := []Person{Person{Name: "Nicolas", Surname: "Contino", Licence: "R-009"}}

	t1s, _ := time.Parse(shortForm, "2020-Feb-03")
	t1e, _ := time.Parse(shortForm, "2020-May-20")
	t2s, _ := time.Parse(shortForm, "2018-Jan-15")
	t2e, _ := time.Parse(shortForm, "2019-Jan-01")
	t3s, _ := time.Parse(shortForm, "2017-Dec-05")
	t3e, _ := time.Parse(shortForm, "2018-Nov-03")

	exps := []VaccineExperiment{
		VaccineExperiment{Company: "International Medical Laboratory", Description: "Testing vaccine for coronavirus",
			Leader: led1, Researchers: res1, VaccineName: "VAC-COV-1", Disease: "COVID-19",
			Conclusions: make([]Conclusion, 0), StartTime: t1s, EndTime: t1e},
		VaccineExperiment{Company: "New York Medical Labolatory", Description: "Testing vaccine for new mutations of common flu.",
			Leader: led2, Researchers: res2, VaccineName: "FLU-42-B", Disease: "Common flu",
			Conclusions: make([]Conclusion, 0), StartTime: t2s, EndTime: t2e},
		VaccineExperiment{Company: "California Main Hospital", Description: "Experimenting with vaccine for well known disease, which is laziness",
			Leader: led3, Researchers: res3, VaccineName: "ANTI-LAZY-v0", Disease: "Laziness",
			Conclusions: make([]Conclusion, 0), StartTime: t3s, EndTime: t3e},
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

func (s *SmartContract) CreateExperiment(ctx contractapi.TransactionContextInterface, experimentKey string, company string, description string,
	leaderName string, leaderSurname string, leaderLicence string, vaccineName string, disease string, startTime string, endTime string) error {

	leader := Person{
		Name:    leaderName,
		Surname: leaderSurname,
		Licence: leaderLicence,
	}

	researchers := []Person{}
	conclusions := []Conclusion{}
	start, err1 := time.Parse(shortForm, startTime)
	end, err2 := time.Parse(shortForm, endTime)

	if err1 != nil {
		return fmt.Errorf("Wrong date: $s", err1.Error())
	}
	if err2 != nil {
		return fmt.Errorf("Wrong date: $s", err2.Error())
	}

	experiment := VaccineExperiment{
		Company:     company,
		Description: description,
		Leader:      leader,
		Researchers: researchers,
		VaccineName: vaccineName,
		Disease:     disease,
		Conclusions: conclusions,
		StartTime:   start,
		EndTime:     end,
	}

	experimentAsBytes, _ := json.Marshal(experiment)

	return ctx.GetStub().PutState(experimentKey, experimentAsBytes)
}

func (s *SmartContract) QueryExperiment(ctx contractapi.TransactionContextInterface, experimentKey string) (*VaccineExperiment, error) {

	experimentAsBytes, err := ctx.GetStub().GetState(experimentKey)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect: %s", err.Error())
	}

	if experimentAsBytes == nil {
		return nil, fmt.Errorf("Experiment %s does not exist", experimentKey)
	}

	experiment := new(VaccineExperiment)
	_ = json.Unmarshal(experimentAsBytes, experiment)

	return experiment, nil
}

func (s *SmartContract) QueryConclusions(ctx contractapi.TransactionContextInterface, experimentKey string) (*QueryConclusionResult, error) {
	experimentBytes, err := ctx.GetStub().GetState(experimentKey)

	if err != nil {
		return nil, fmt.Errorf("Failed to connect: %s", err.Error())
	}

	if experimentBytes == nil {
		return nil, fmt.Errorf("Experiment %s does not exist", experimentKey)
	}

	experiment := new(VaccineExperiment)
	_ = json.Unmarshal(experimentBytes, experiment)

	result := QueryConclusionResult{
		Key:     experimentKey,
		Records: experiment.Conclusions,
	}

	return &result, nil
}

func (s *SmartContract) AddConclusion(
	ctx contractapi.TransactionContextInterface,
	experimentKey string,
	authorName string,
	authorSurname string,
	authorLicence string,
	content string) error {

	experimentBytes, err := ctx.GetStub().GetState(experimentKey)

	if err != nil {
		return err
	}

	if experimentBytes == nil {
		return fmt.Errorf("Experiment %s does not exist", experimentKey)
	}

	experiment := new(VaccineExperiment)
	_ = json.Unmarshal(experimentBytes, experiment)

	author := Person{
		Name:    authorName,
		Surname: authorSurname,
		Licence: authorLicence,
	}

	conclusion := Conclusion{
		Author:       author,
		Content:      content,
		CreationTime: time.Now(),
	}

	experiment.Conclusions = append(experiment.Conclusions, conclusion)

	updatedExperimentBytes, _ := json.Marshal(experiment)

	return ctx.GetStub().PutState(experimentKey, updatedExperimentBytes)
}

func (s *SmartContract) AddResearcherToExperiment(
	ctx contractapi.TransactionContextInterface,
	experimentKey string,
	name string,
	surname string,
	licence string) error {

	experimentBytes, err := ctx.GetStub().GetState(experimentKey)

	if err != nil {
		return err
	}

	if experimentBytes == nil {
		return fmt.Errorf("Experiment %s does not exist", experimentKey)
	}

	experiment := new(VaccineExperiment)
	_ = json.Unmarshal(experimentBytes, experiment)

	researcher := Person{
		Name:    name,
		Surname: surname,
		Licence: licence,
	}

	experiment.Researchers = append(experiment.Researchers, researcher)

	updatedExperimentBytes, _ := json.Marshal(experiment)

	return ctx.GetStub().PutState(experimentKey, updatedExperimentBytes)
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
