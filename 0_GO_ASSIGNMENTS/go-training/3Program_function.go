package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func substract(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func divide(a, b float32) float32 {

	if b == 0 {
		fmt.Println("Value cannot divide by zero ")
		return 0
	}
	return a / b
}

func main() {
	fmt.Println("Addition of two numbers :", add(10, 30))
	fmt.Println("substraction of two numbers :", substract(10, 30))
	fmt.Println("multiplication of two numbers :", multiplication(10, 30))
	fmt.Println("dividation of two numbers :", divide(10, 0))
}
