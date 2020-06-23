package main

import "fmt"

type billing struct {
	m map[string]float64
}

// all users have o balance
// balance can be adjusted: subscriberid,amount, return: bool,  string
// balance can be retrieved: subscriberid, return: float64
// amount can be transfered from one acc to another: sibscriberAid, subscriberBid, amount, return: bool, string
//

func main() {

	b := newBilling()

	fmt.Println(b.retrieve("Gunel"))

	fmt.Println(b.adjust("Gunel", -20))

	fmt.Println(b.retrieve("Gunel"))

	fmt.Println(b.adjust("Gunel", 10))

	fmt.Println(b.retrieve("Gunel"))
	fmt.Println(b.retrieve("Kanan"))

	fmt.Println(b.transfer("Gunel", "Kanan", -20))

	fmt.Println(b.retrieve("Gunel"))
	fmt.Println(b.retrieve("Kanan"))

	fmt.Println(b.transfer("Gunel", "Kanan", 20))

	fmt.Println(b.retrieve("Gunel"))
	fmt.Println(b.retrieve("Kanan"))

	fmt.Println(b.transfer("Gunel", "Kanan", 7))

	fmt.Println(b.retrieve("Gunel"))
	fmt.Println(b.retrieve("Kanan"))
}

func newBilling() *billing {
	return &billing{
		m: map[string]float64{},
	}
}

func (b *billing) adjust(id string, amount float64) (bool, string) {
	if b.m[id]+amount < 0 {
		return false, "Balance can not be negative"
	}
	b.m[id] += amount
	return true, ""
}

func (b *billing) retrieve(id string) (string, float64) {
	return id, b.m[id]
}

func (b *billing) transfer(idA string, idB string, amount float64) (bool, string) {
	if b.m[idB]+amount < 0 {
		return false, "Balance B can not be negative"
	}
	if b.m[idA]-amount < 0 {
		return false, "Transfer amount exceeds SubscriberA  balance"
	}
	b.m[idA] -= amount
	b.m[idB] += amount
	return true, ""
}
