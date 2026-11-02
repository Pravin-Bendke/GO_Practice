package main

import "fmt"

func main() {

	employees := make(map[string]int) //store data in map

	employees["Pravin"] = 23 //Add Data
	employees["Rahul"] = 24  //Add Data
	employees["Sagar"] = 34
	employees["Srushti"] = 25

	fmt.Println("Employee Details")
	for name, age := range employees {
		fmt.Println("Name:", name, "| Age:", age)
	}

	age, is_Exists := employees["Pravin"] //checking Key is Available or Not

	if is_Exists {
		fmt.Println("Pravin is Found :", age, is_Exists)
	} else {
		fmt.Println("Record not found!...")
	}

	delete(employees, "Pravin") //Delete Employee
	fmt.Println("After deleting record")

	for name, age := range employees {
		fmt.Println(name, "|", age)
	}
}
