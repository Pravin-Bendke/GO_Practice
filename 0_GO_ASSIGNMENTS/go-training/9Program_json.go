package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	ID         int    `json :"id"` // json struct tag  #It tells Go to use id as the JSON field name instead of ID.
	Name       string `json :"name"`
	Email      string `json :"email"`
	Department string `json :"department"`
}

func main() {

	employee := Employee{
		ID:         1,
		Name:       "Pravin Bendke",
		Email:      "pravin@gmail.com",
		Department: "IT Department",
	}
	jsonData, err := json.MarshalIndent(employee, "", " ")
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("! STRUCT TO JSON !")
	fmt.Print(string(jsonData))

	var newEmployee Employee
	err = json.Unmarshal(jsonData, &newEmployee)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println("\n! JSON TO STRUCT !")
	fmt.Println(newEmployee)
}
