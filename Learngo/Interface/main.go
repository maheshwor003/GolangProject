package main

import "fmt"

// Define an interface
type LaptopRepair interface {
	Repair() string
}
type geometry interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}
type rect struct {
	lenght, breath float64
}

func (r rect) area() float64 {
	return r.breath * r.lenght
}
func (r rect) perimeter() float64 {
	return r.breath + r.lenght
}
func (s square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return s.side + s.side
}

func Shop(repairer LaptopRepair) string {
	return repairer.Repair()
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

// Define a struct that implements the interface
//Code lai less coupling banauna ani depencey kam garna

type Engineer struct {
	Name        string
	Proffession string
	Department  string
}

type Farmer struct {
	Name     string
	LandArea int
}

func (e Engineer) Repair() string {
	return fmt.Sprintf("Engineer %s repaired the laptop", e.Name)

}
func (e Farmer) Repair() string {
	return fmt.Sprintf("Farmer  %s repaired the laptop", e.Name)

}

func main() {

	engineer1 := Engineer{
		Name:        "John Doe",
		Proffession: "Software Engineer",
		Department:  "IT",
	}

	farmer1 := Farmer{
		Name:     "Jane Smith",
		LandArea: 100,
	}

	status := Shop(engineer1)
	fmt.Println(status)
	status2 := Shop(farmer1)
	fmt.Println(status2)

	s := square{side: 10.0}
	fmt.Println("Area of Square", s.area())
	measure(s)

	r := rect{lenght: 2, breath: 5}
	fmt.Println("Area of Rectangle", r.area())
	fmt.Println("Perimeter of Rectangle", r.perimeter())
	measure(r)

}
