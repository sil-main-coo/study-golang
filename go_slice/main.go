package main

import "fmt"

func main() {
	// slice_name := []datatype{values}
	slice := []int{}
	slice1 := []int{1, 2, 3}

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Create a Slice From an Array
	// var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array

	arr := [6]int{10, 11, 12, 13, 14, 15}
	slice3 := arr[0:4]
	fmt.Printf("slice3 = %v\n", slice3)
	fmt.Printf("length = %d\n", len(slice3))
	fmt.Printf("capacity = %v\n", cap(slice3))
}
