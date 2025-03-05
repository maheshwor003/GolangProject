package main

import (
	"fmt"
)

func main() {

	//Primite data types

	// int,float32,string,rune  alias of int 32,slice ,map, struct,interface bool

	// a := 10
	// var b float32 = 3.5
	// var z string = "this is nepal"
	// var s bool = false
	// word := "hello" //can be make like this //declare and assign

	// fmt.Print(a)
	// fmt.Print(b)
	// fmt.Print(z)
	// fmt.Print(s)

	// fmt.Print(word)

	// fmt.Print("Hello World")

	//conditional statment in golang
	// we have if and switch
	age := 20

	if age > 18 {
		fmt.Println("Eligible to Vote")
	} else {
		fmt.Println("Not Eligible to Vote")
	}

	switch {
	case age < 18:
		fmt.Println("Eligible to Vote")
	case age > 18:
		fmt.Println("Not Eligible to Vote")
	}

	//Slices  and Array
	//list dynamic increase size

	var names []string = []string{}
	names = append(names, "Maheshwor")
	names = append(names, "Shakha")
	names = append(names, "okay")

	fmt.Println(names)
	fmt.Println(len(names))

	//looping

	for index, value := range names {
		fmt.Println("index:", index, " value: ", value)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	var numbers []int = []int{1, 2, 3, 5}
	for ind := range numbers {

		numbers[ind] += 1
	}
	fmt.Println(numbers)

}
