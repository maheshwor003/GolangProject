package main

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(f string) {
	fmt.Println("Reading File")
	data, err := os.ReadFile(f)
	if err != nil {
		log.Fatal("Error in reading   a File")
	}

	fmt.Println(string(data))

}
func WriteFile(f string) {
	file, err := os.Create(f)
	if err != nil {
		log.Fatal("Error Ceating  a File")
	}
	len, err := file.WriteString("manuallly right data and test ")
	if err != nil {
		log.Fatal("Error writing  a File")
	}
	defer file.Close()
	fmt.Println("File Writed Successfully", file.Name())
	fmt.Println("File  Length ", len)

}
func CreateFile(f string) {
	file, err := os.Create(f)
	if err != nil {
		log.Fatal("Error creating a File")

	}
	defer file.Close()
	fmt.Println("File Created Successfully")

}

func DeleteFile(f string) {
	fmt.Println("Delete File")
	err := os.Remove(f)
	if err != nil {
		log.Fatal("Error in Deleting    a File")
	}
	fmt.Println("File Deleted Successfully")

}
func main() {

	fmt.Println("handling")
	fmt.Println("test file handling ")

	// //Creating a file
	CreateFile("Maheshwor.txt")

	// //Writing a file
	WriteFile("Maheshwor.txt")

	//Read a File
	ReadFile("Maheshwor.txt")

	DeleteFile("Maheshwor.txt")

}
