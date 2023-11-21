package main

import "fmt"

func main() {
	a := 10.2

	var pointer *float64

	pointer = &a

	fmt.Println(pointer) // print 'a' address
	fmt.Printf("%T\n", pointer)

	b := 123
	p2 := new(int)
	p2 = &b

	fmt.Println(p2) // print 'b' address
	fmt.Printf("%T\n", p2)

	// ZERO VALUE OF POINTER
	var p3 *int
	p4 := new(int)

	fmt.Println(p3) // is nil
	fmt.Println(p4) // is not nil

	///

	c := 111
	var p5 *int = &c

	fmt.Println(c) // 111
	*p5 = 999
	fmt.Println(c) // 999

	// POINTER OF ARRAY

	arr := [3]int{1, 2, 3}
	var p6 *[3]int = &arr
	fmt.Println(p6) //
	fmt.Printf("%T\n", p6)

	d := 222
	fmt.Println(d) // 222
	applyPointer(&d)
	fmt.Println(d) // 777
}

func applyPointer(pointer *int) {
	*pointer = 777
}
