package main

import "fmt"

// with Generic
func printSlice[T int | string | bool](items []T) {
	// func printSlice[T interface{}](items []T) {
	// func printSlice[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

//Wrong Way Duplicate
//without generic

// func printStringSlice(items []string) {
// 	for _, item := range items {
// 		fmt.Println(item)
// 	}
// }

//LIFO
//we are using in struct

type stack[T any] struct {
	element []T
}

func main() {

	myStack := stack[string]{
		// func printSlice[T interface{}](items []T) {
		// element: []int{1, 2, 3},
		element: []string{"Pravin", "Rahul"},
	}

	fmt.Println(myStack)

	// nums := []int{1, 2, 3, 4, 3, 3}
	// names := []string{"Pravin", "Rahul", "Srushti"}
	// vals := []bool{true, false, true}
	// printSlice(vals)

}
