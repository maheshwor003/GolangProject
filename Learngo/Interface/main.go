package main

import "fmt"

// Define an interface
type LaptopRepair interface {
	Repair() string
}

func Shop(repairer LaptopRepair) string {
	return repairer.Repair()
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
}
