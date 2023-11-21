package main

import "fmt"

const PI = 3.14

const (
	A int = 1
	B     = 1.23
	C     = "Hi!"
)

func main() {
	const PI1 = 3.14

	var number int

	number = 10

	fmt.Println(number)

	var number1 int = 11

	fmt.Println(number1)

	// Type Inference
	var email = "abc@gmail.com"

	fmt.Println(email)

	var (
		name    string
		address string
		age     int
	)

	name = "David"
	address = "Vietnam"
	age = 25

	fmt.Println(name)

	fmt.Println(address)

	fmt.Println(age)

	// PI = 1 --> fail
	fmt.Println(PI)
	fmt.Println(PI1)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

}
 