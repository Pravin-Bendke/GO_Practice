package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Employee struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Dept  string `json:"dept"`
	MobNo int    `json:"mobNo"`
}

func main() {
	employees := []Employee{
		{
			ID:    1,
			Name:  "Pravin",
			Email: "pravin@gmail.com",
			Dept:  "IT Dept",
			MobNo: 9767196869,
		},
		{
			ID:    2,
			Name:  "rahul",
			Email: "rahul@gmail.com",
			Dept:  "IT Dept",
			MobNo: 1234567890,
		},
		{
			ID:    3,
			Name:  "srushti",
			Email: "srushti@gmail.com",
			Dept:  "testing Dept",
			MobNo: 9876543210,
		},
	}
	jsonData, err := json.MarshalIndent(employees, "", " ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	err = os.WriteFile("employee.json", jsonData, 0644)
	if err != nil {
		fmt.Println("file ahndling Error", err)
		return
	}
	fmt.Println("Data Written Successfull !")
	fileData, err := os.ReadFile("employee.json")
	if err != nil {
		fmt.Println("file read Error ", err)
		return
	}
	var newEmployees []Employee
	err = json.Unmarshal(fileData, &newEmployees)
	if err != nil {
		fmt.Println("Data Decode Error", err)
		return
	}
	fmt.Println("Data Retrive Successfully !")
	for _, employee := range newEmployees {
		fmt.Println("--------------------------------------------------")

		fmt.Println("ID :", employee.ID)
		fmt.Println("Name :", employee.Name)
		fmt.Println("Dept :", employee.Dept)
		fmt.Println("Mobile Number :", employee.MobNo)
		fmt.Println("")
	}

}
