package main

import "fmt"

func modify(x *int) {
	*x = 10
}

func modifyStruct(x *Employee) {
	x.age = 28
}

func modifyArray(x []int) {
	x[2] = 28
}

type Employee struct {
	name string
	age  int
}

func main() {

	// Pointer is a variable that stores the memory address of another variable
	a := 3
	var ptr *int = &a
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(*ptr)

	// Pointer to function
	x := 3
	fmt.Println(x)
	modify(&x)
	fmt.Println(x)

	// Pointer to data structure
	s := Employee{name: "Maheshwor", age: 27}
	fmt.Println(s)
	modifyStruct(&s)
	fmt.Println(s)

	// Using new to allocate memory
	ptr2 := new(int) // Renamed to avoid conflict
	*ptr2 = 42
	fmt.Println(*ptr2)

	r := make([]int, 3)
	r[0], r[1], r[2] = 1, 2, 4
	fmt.Println(r)
	modifyArray(r)
	fmt.Println(r)

}
