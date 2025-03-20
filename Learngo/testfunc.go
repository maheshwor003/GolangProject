package main

import (
	"fmt"
)

type Student struct {
	name       string
	department string
	age        int
	grade      int
	height     float32
	College
}

type College struct {
	CollegeName        string
	CollegeLocation    string
	CollegePhoneNumber int64
}

// Method
func (s Student) GetStudentInfo() {
	fmt.Print(s.name, "is reading")
}

func (s Student) CopyStudentInfo(Student Student) {
	fmt.Print(s.name, "is Copying ", Student.name)
}

func (a College) GetCollegeName() string {
	return a.CollegeName
}

func NewStudent(name string, Department string, age int, grade int, height float32,
	collegename string, collegelocation string, collegephonenumber int64) Student {
	student := Student{
		name:       name,
		department: Department,
		age:        age,
		grade:      grade,
		College: College{
			CollegeName:        collegename,
			CollegeLocation:    collegelocation,
			CollegePhoneNumber: collegephonenumber,
		},
	}
	return student

}

func AgeStudent(name string, Department string, age int, grade int, height float32) int {
	student := Student{
		name:       name,
		department: Department,
		age:        age,
		grade:      grade,
	}
	return student.age

}

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

	txt1 := "Hello!"
	txt2 := ""
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

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

	//struct in go lang
	// like  class

	var student1 Student = Student{
		name:       "Maheshwor Shakha",
		department: "Mechanical Engineeering",
		grade:      2,
		age:        23,
		height:     22.3,
	}

	fmt.Println(student1)
	fmt.Println(student1.name)

	student11 := NewStudent("ram", "computer", 22, 32, 33.3, "NEC college", "Changunaryan", 34234324)
	fmt.Println(student11)

	fmt.Println(student11.CollegeLocation)

	fmt.Println(student11.College.CollegeLocation)

	fmt.Println(student11.GetCollegeName())

	ages := AgeStudent("ram", "computer", 27, 40, 33.3)
	fmt.Println(ages)

	student1.GetStudentInfo()

	student11.CopyStudentInfo(student1)

	student1.GetStudentInfo()

}
