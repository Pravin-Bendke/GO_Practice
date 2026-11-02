package main

import "fmt"

func main() {
	//with var
	var name string = "Pravin Bendke"
	var addr string = "Vadner Bhairav"
	var mob int = 9767196869
	var is_Employee bool = true
	var percentage float64 = 60.77

	//with const

	const collegeName string = "CMCS College Nashik Ganagapur Road"
	const collegeCode int = 818
	const companyName string = "Veritas GXP"

	//With :=
	department := "Development"
	c_Addr := "Ambad Nashik Cidco"

	//student
	fmt.Println("name : ", name)
	fmt.Println("addr : ", addr)
	fmt.Println("mob : ", mob)
	fmt.Println("is_Employee : ", is_Employee)
	fmt.Println("percentage : ", percentage)

	//college/company Information
	fmt.Println("collegeName :", collegeName)
	fmt.Println("collegeCode :", collegeCode)
	fmt.Println("companyName :", companyName)

	fmt.Println("department :", department)
	fmt.Println("c_Addr :", c_Addr)

}
