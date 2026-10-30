package main

import (
	"fmt"
	"time"
)

// type student struct {
// 	Name      string
// 	age       int
// 	marks     float64
// 	mobNo     int64
// 	createdAt time.Time
// }

// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time //nano seconds count
// }

// func newOrder(id string, amount float32, status string) *order {
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}
// 	return &myOrder
// }

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o *order) getAmount() float32 {
// 	return o.amount
// }

//struct Embeding

type customer struct {
	name  string
	phone string
	addr  string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time //nano seconds count
	customer
}

func main() {

	// myOrder := newOrder("1", 350.00, "received")
	// fmt.Println(myOrder)

	// myOrder := order{
	// 	id:     "1",
	// 	amount: 100.00,
	// 	status: "Received",
	// 	createdAt: time.Now(),
	// }
	// myOrder.status = "Paid"
	// fmt.Println("Order Details:", myOrder)
	// myOrder.changeStatus("completed")
	// myOrder.createdAt = time.Now()

	// fmt.Println("Order Details:", myOrder)
	// fmt.Println("Order Amount Details:", myOrder.getAmount())

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    300.00,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// fmt.Println("Order Details:", myOrder2)
	// fmt.Println("Order Details:", myOrder)

	// s1 := student{
	// 	Name:      "Pravin",
	// 	age:       25,
	// 	marks:     60,
	// 	mobNo:     1234567890,
	// 	createdAt: time.Now(),
	// }

	// fmt.Println("Information of", s1.Name, ":", s1)

	//struct embedding

	C1 := customer{
		name:  "Pravin",
		phone: "9767196869",
		addr:  "Vadner Bhairav",
	}

	O1 := order{
		id:        "101",
		amount:    2000.00,
		status:    "Packed",
		createdAt: time.Now(),
		customer:  C1,
	}
	fmt.Println(O1)

}
