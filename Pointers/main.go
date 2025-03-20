package main

import "fmt"

// Function that modifies the value and returns the updated value
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
}
