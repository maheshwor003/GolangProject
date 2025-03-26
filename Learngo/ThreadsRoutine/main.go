// Concurrency in Go
// Concurrency in Go is achieved by using goroutines.
//  A goroutine is a lightweight thread managed by the Go runtime.
//   It is a function that runs concurrently with other functions. G
//   oroutines are created by using the go keyword followed by a function invocation.
//    The following example demonstrates how to create a goroutine:

package main

import (
	"fmt"
	"sync"
	"time"
)

func printHello(wg *sync.WaitGroup) {
	defer wg.Done() //this method will call at the end of logic
	fmt.Println("Hello")
}

func printHelloAgain(wg *sync.WaitGroup) {
	defer wg.Done() //this method will call at the end of logic
	fmt.Println("Hello world ")
}

func connecttodatabase() (bool, error) {
	fmt.Println("Connecting to the database")
	time.Sleep(2 * time.Second)
	fmt.Println("Connected to the database")
	return true, nil
}

func ConnectThirdPartyAPI() (bool, error) {
	fmt.Println("Connecting to the third party API")
	time.Sleep(2 * time.Second)
	fmt.Println("Connected to the third party API")
	return true, nil
}

func main() {

	//multi threading
	// printHello()
	// go printHelloAgain()
	// time.Sleep(3 * time.Second)

	var wg sync.WaitGroup //for waiting for go routines to execute
	wg.Add(2)             //to tell how many go routines are there

	go printHelloAgain(&wg)
	go printHello(&wg)
	wg.Wait() // to tell main function to wait unitl go routines are done with  execution

	ws := sync.WaitGroup{}
	timeStart := time.Now()
	ws.Add(1)
	go func() {
		status, _ := connecttodatabase()
		fmt.Println("Status ", status)
		ws.Done()
	}()
	ws.Add(1)
	go func() {

		status, _ := connecttodatabase()
		fmt.Println("Status ", status)
		wg.Done()
	}()
	ws.Add(1)
	go func() {
		status, _ := connecttodatabase()
		fmt.Println("Status ", status)
		ws.Done()
	}()

	// if err != nil {
	// 	fmt.Printf("Error connecting to the database: %v\n", err)
	// 	return
	// }

	status, err := connecttodatabase()
	if err != nil {
		fmt.Println("error connecting to the database")
		return
	}

	if !status {
		fmt.Println("database connection failed")
		return
	}

	//connect to third party API

	thirdpartystatus, _ := ConnectThirdPartyAPI()
	// if thirdpartyerror != nil {
	// 	fmt.Println("error connecting to the third party API")
	// 	return
	// }

	if !thirdpartystatus {
		fmt.Println("third party API connection failed")
		return
	}

	fmt.Println("All Services connected successfully")
	ws.Wait()
	elapsed := time.Since(timeStart)
	fmt.Printf("Time taken: %s\n", elapsed)
}
