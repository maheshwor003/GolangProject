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

	wg := sync.WaitGroup{}

	timeStart := time.Now()
	wg.Add(1)
	go func() {
		status, _ := connecttodatabase()
		fmt.Println("Status ", status)
		wg.Done()
	}()
	wg.Add(1)
	go func() {

		status, _ := connecttodatabase()
		fmt.Println("Status ", status)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		status, _ := connecttodatabase()
		fmt.Println("Status ", status)
		wg.Done()
	}()

	// if err != nil {
	// 	fmt.Println("error connecting to the database")
	// 	return
	// }

	// status, err := connecttodatabase()
	// if err != nil {
	// 	fmt.Println("error connecting to the database")
	// 	return
	// }

	// if !status {
	// 	fmt.Println("database connection failed")
	// 	return
	// }

	// //connect to third party API

	// thirdpartystatus, _ := ConnectThirdPartyAPI()
	// // if thirdpartyerror != nil {
	// // 	fmt.Println("error connecting to the third party API")
	// // 	return
	// // }

	// if !thirdpartystatus {
	// 	fmt.Println("third party API connection failed")
	// 	return
	// }

	// fmt.Println("All Services connected successfully")
	wg.Wait()
	elapsed := time.Since(timeStart)
	fmt.Printf("Time taken: %s\n", elapsed)
}
