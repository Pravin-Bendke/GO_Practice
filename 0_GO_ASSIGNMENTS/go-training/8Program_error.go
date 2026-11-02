package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID         int
	Name       string
	Email      string
	Department string
}

var employees = []Employee{
	{1, "Pravin", "pravin@gmail.com", "it"},
	{2, "rahul", "rahul@gmail.com", "it"},
	{3, "srushti", "srushti@gmail.com", "it"},
	{4, "sakshi", "sakshi@gmail.com", "it"},
}

func getEmployeeDetails(id int) (*Employee, error) {
	for detail := range employees {
		if employees[detail].ID == id {
			return &employees[detail], nil
		}
	}
	return nil, errors.New("Employee Not Found")
}

func main() {
	employee, err := getEmployeeDetails(1)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Employee Found", *employee)
	}

	employee, err = getEmployeeDetails(10)
	if err != nil {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Employee Found", *employee)
	}

}
