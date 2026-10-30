package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("inside changeNum function", num)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println("inside of main", num)
}

// package main

// import "fmt"

// type Student struct {
// 	name string
// }

// func (n *Student) greet() string {
// 	n.name = "jadhav"
// 	return n.name
// }

// func main() {
// 	student1 := Student{name: "Rahul"}
// 	student1.greet()
// 	fmt.Println(student1.name)

// }
