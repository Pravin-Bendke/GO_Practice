// package main

// import "fmt"

// type OrderStatus int

// const (
// 	received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delevered
// )

// func changeOrderStatus(status OrderStatus) {
// 	fmt.Println("current order status is", status)
// }

// func main() {
// 	changeOrderStatus(Prepared)
// }

package main

import "fmt"

type OrderStatus string

const (
	received  OrderStatus = "received"
	Confirmed             = "Confirmed"
	Prepared              = "Prepared"
	Delevered             = "Delevered"
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Current status is", status)
}

func main() {
	changeOrderStatus(received)
}
