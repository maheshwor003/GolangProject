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
	StudentNumber      int
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
	collegename string, collegelocation string, collegephonenumber int64, studentumber int) Student {
	student := Student{
		name:       name,
		department: Department,
		age:        age,
		grade:      grade,
		height:     height,
		College: College{
			CollegeName:        collegename,
			CollegeLocation:    collegelocation,
			CollegePhoneNumber: collegephonenumber,
			StudentNumber:      studentumber,
		},
	}
	return student

}

func main() {

	colleges := []College{
		{CollegeName: "NEC College", CollegeLocation: "", CollegePhoneNumber: 0, StudentNumber: 322},
		{CollegeName: "KEC College", CollegeLocation: "", CollegePhoneNumber: 0, StudentNumber: 12},
		{CollegeName: "Pulchoki College", CollegeLocation: "", CollegePhoneNumber: 0, StudentNumber: 322},
		{CollegeName: "Cosmos College", CollegeLocation: "", CollegePhoneNumber: 0, StudentNumber: 121},
	}

	// Initialize map
	//var CollegeStudentNumber = map[string]int{}
	CollegeStudentNumber := make(map[string]int)
	for _, college := range colleges {
		CollegeStudentNumber[college.CollegeName] = college.StudentNumber
	}

	studentInNEC, hasValue := CollegeStudentNumber[" College"]
	if !hasValue {
		fmt.Println("NEC College is not in the list")
	} else {
		fmt.Println("Number of students in NEC College:", studentInNEC)
	}

	studentInNEC = CollegeStudentNumber["KEC College"]
	fmt.Println("Number of students in KEC College:", studentInNEC)
}
