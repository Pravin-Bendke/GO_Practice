package main

import "fmt"

type Payment interface{ Pay(amount float64) }

type Cash struct{}

func (c Cash) Pay(amount float64) {
	fmt.Println("Paid Rs.", amount, "using cash")
}

type Upi struct{}

func (u Upi) Pay(amount float64) {
	fmt.Println("Paid Rs.", amount, "using Upi")
}

type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) {
	fmt.Println("Paid Rs.", amount, "using CreditCard")
}

func paymentProcess(p Payment, amount float64) {
	p.Pay(amount)
}

func main() {
	cash := Cash{}
	upi := Upi{}
	creditcard := CreditCard{}

	paymentProcess(cash, 2000)
	paymentProcess(upi, 2000)
	paymentProcess(creditcard, 2000)

}
