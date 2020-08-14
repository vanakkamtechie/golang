package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(8)
	fmt.Println(runtime.NumCPU())
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Starting Go Routines")
	go func() {
		defer wg.Done()

		time.Sleep(1 * time.Microsecond)
		for char := 'a'; char < 'a'+26; char++ {
			fmt.Printf("%c ", char)
		}
	}()

	fmt.Println("Started Go Routine 1")

	go func() {
		defer wg.Done()

		for number := 1; number < 27; number++ {
			fmt.Printf("%d ", number)
		}
	}()

	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
