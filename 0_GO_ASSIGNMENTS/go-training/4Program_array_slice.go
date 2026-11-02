package main

import (
	"fmt"
)

func addEmployee(employee []string, name string) []string {
	employee = append(employee, name)
	return employee
}

func removeEmployee(employees []string, name string) []string {
	for i, employee := range employees {
		if employee == name {
			employees = append(employees[:i], employees[i+1:]...)
			return employees
		}
	}
	fmt.Println("not fount")
	return employees
}

func displayData(employees []string) {
	fmt.Println("!! Employee Data !!")
	for i, emp := range employees {
		fmt.Print(i+1, ". ", emp, "\n")
	}
}

func searchEmployee(employees []string, name string) {
	for _, emp := range employees {
		if emp == name {
			fmt.Println("Employee Found :", emp)
			return
		}
	}
	fmt.Println("Employee Not Found :", name)
}

func main() {
	employees := []string{"Pravin", "Rahul", "Srushti", "Sachin"}
	displayData(employees)
	employee := removeEmployee(employees, "Pravin") //Remove Data
	displayData(employee)                           // Display data
	searchEmployee(employees, "Pravin")             //Searching Data "Pravin"
	searchEmployee(employees, "Pradip")
	displayData(employees) //Display Data

}
