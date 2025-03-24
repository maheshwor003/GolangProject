package main

import "fmt"

type Student struct {
	name string
	age  int
}

//Receiver function
func (s *Student) update(age int) {
	s.age = age
}

// Function that modifies the value and returns the updated value

func updateStudent(s *Student, age int) {
	s.age = age // Return the updated value
}

func updateValue(num *int) int {
	fmt.Println("Inside Function - Address:", num) // Print address inside function
	*num = *num + 10                               // Modify the original value
	return *num                                    // Return the updated value
}

func main() {
	value := 20
	fmt.Println("Before - Value:", value, "Address:", &value) // Print initial address

	newValue := updateValue(&value) // Passing the address of value

	fmt.Println("After - Value:", value, "Address:", &value) // Print modified value and address
	fmt.Println("Returned Value:", newValue)                 // Printing the returned value

	sr := Student{name: "Maheshwor", age: 13}
	updateStudent(&sr, 12)
	fmt.Println(sr)

	s := Student{name: "Ram", age: 11}
	s.update(21)
	fmt.Println(s)

	//map techniques 1
	var student map[int]string
	student = map[int]string{
		1: "Maheshwor",
	}
	fmt.Println(student)

	//map techinques 2

	// Correctly initializing a map
	studentAge := make(map[string]int)

	// Correctly adding key-value pairs to the map
	studentAge["Mahesh"] = 23
	studentAge["Rama"] = 22

	val, ok := studentAge["Mahesh"] // here ok represent true  boolean

	fmt.Println(val, ok)
	studentAge["Mahesh"] = 27
	va, ok := studentAge["Mahesh"]
	fmt.Println(va, ok)
	// Printing the map
	fmt.Println(studentAge)

	delete(studentAge, "Mahesh")
}
