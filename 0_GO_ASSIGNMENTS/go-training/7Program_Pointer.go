package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Department string
}

func updateEmployee(employee *Employee) {
	employee.ID = 1
	employee.Name = "Pravin Bendke"
	employee.Email = "pravin@gmail.com"
	employee.Department = "IT Department"
}

func main() {

	emp := Employee{
		ID:         1,
		Name:       "Rahul Jadhav",
		Email:      "rahul@gmail.com",
		Department: "IT Department",
	}
	fmt.Println("!! Before Update Details !!")
	fmt.Println(emp)

	updateEmployee(&emp) //Update employee information
	fmt.Println("!! After Update details !!")

	fmt.Println(emp)
}
