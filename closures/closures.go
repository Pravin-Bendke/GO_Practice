package main

import "fmt"

func counter() (func() int, func() int) {
	count := 0
	increment := func() int {
		count += 1
		return count
	}
	decrement := func() int {
		count -= 1
		return count
	}
	return increment, decrement
}

func main() {
	increment, decrement := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(decrement())
	fmt.Println(decrement())
}
