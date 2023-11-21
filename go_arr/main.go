package main

import "fmt"

func main() {

	// var array_name = [length]datatype{values} // here length is defined

	// or

	// var array_name = [...]datatype{values} // here length is inferred

	var arr = [3]int{1, 2, 3}
	var arr1 = [...]int{4, 5, 6, 7}

	fmt.Println(arr)
	fmt.Println(arr1)

	// array_name := [length]datatype{values} // here length is defined

	// or

	// array_name := [...]datatype{values} // here length is inferred

	arr3 := [3]int{8, 9, 10}
	arr4 := [...]int{11, 12, 13}

	fmt.Println(arr3)
	fmt.Println(arr4)

	arr[0] = -99
	fmt.Println(arr)

	arr5 := [5]int{1: 10, 2: 40}
	fmt.Println(arr5)

	arr6 := [...]string{"abc", "xyz", "ijk"}
	fmt.Println(len(arr5))
	fmt.Println(len(arr6))
}
