package main

import "fmt"

type StudentInfo struct {
	class string
	email string
	age   int
}

// Nested struct
type Student struct {
	id   int
	name string
	info StudentInfo
}

func main() {
	// Named
	st1 := Student{
		id:   123,
		name: "David John",
		info: StudentInfo{
			class: "12A",
			email: "acb@gmail.com",
			age:   17,
		},
	}

	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)
	fmt.Println(st1.info)

	// Normal
	st2 := Student{456, "Robin Hood", StudentInfo{"12A", "acb@gmail.com", 17}}
	fmt.Println(st2)
	fmt.Println(st2.id)
	fmt.Println(st2.name)
	fmt.Println(st2.info)

	//
	// var st3 Student = struct {
	// 	id   int
	// 	name string
	// }{
	// 	id:   999,
	// 	name: "Bill",
	// }
	// fmt.Println(st3)
	// fmt.Println(st3.id)
	// fmt.Println(st3.name)

	// anonymous struct
	var anonymous = struct {
		email string
		age   int
	}{
		email: "abc@gmail.com",
		age:   12,
	}

	fmt.Println(anonymous)

	// struct pointer
	pointer := &Student{
		id:   999,
		name: "Robin",
		info: StudentInfo{
			class: "12A",
			email: "acb@gmail.com",
			age:   17,
		},
	}

	fmt.Println(&pointer)
	fmt.Println(pointer.id)
	fmt.Println(pointer.name)
	fmt.Println(pointer.info)

	// anonymous field
	type NoName struct {
		string
		int
	}

	n := NoName{"Truong Viet Hoang", 24}
	fmt.Println(n)

	// Compare struct
	type struc1 struct {
		id   int
		name string
		// info map[int]int // un comment
	}

	s1 := struc1{1, "A"} // map[int]int{0: 1}

	s2 := struc1{1, "A"} // map[int]int{0: 1}

	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}

	// Zero value
	var student Student
	fmt.Println(student)
	student.id = 1
	fmt.Println(student)
}
