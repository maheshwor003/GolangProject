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
		height:     height,
		College: College{
			CollegeName:        collegename,
			CollegeLocation:    collegelocation,
			CollegePhoneNumber: collegephonenumber,
		},
	}
	return student

}

func main() {

}
