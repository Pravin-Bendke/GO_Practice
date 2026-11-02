package main

import "fmt"

type Employee struct {
	ID         int
	Name       string
	Email      string
	Department string
}

func main() {
	employee1 := Employee{
		ID:         01,
		Name:       "Pravin",
		Email:      "pravinbendke0@gmail.",
		Department: "IT Department",
	}

	fmt.Println("!! Employee Details !!")
	fmt.Println("ID :", employee1.ID)
	fmt.Println("Name:", employee1.Name)
	fmt.Println("Email:", employee1.Email)
	fmt.Println("Department:", employee1.Department)

}
