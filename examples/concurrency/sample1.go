package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	ch1 <- 1
	wg.Add(1)
	go showOddNumbers(ch1, ch2, &wg)
	wg.Add(1)
	go showEvenNumbers(ch1, ch2, &wg)

	wg.Wait()
	fmt.Println("bye")
	//close(ch1)
	//close(ch2)
}

func showOddNumbers(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {

	for {
		time.Sleep(time.Millisecond * 1000)
		val := <-ch1

		//_, isOpen := <-ch2
		//fmt.Println("Channel 2 status:", isOpen)
		//if isOpen {

		//}
		if val >= 10 {
			defer func() {
				fmt.Println("Wg done for odd numbers")
				wg.Done()
			}()
			fmt.Println("Channel 1 has been closed")
			close(ch1)
			goto skipFor
			//break
		} else {
			fmt.Println("Odd Number is:", val)
			ch2 <- val + 1
		}

	}

skipFor:
	fmt.Println("Skipped For loop showOddNumbers")
}

func showEvenNumbers(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	var val int
	for {
		time.Sleep(time.Millisecond * 1000)
		val = <-ch2

		if val >= 10 {
			defer func() {
				fmt.Println("Wg done for even numbers")
				wg.Done()
			}()
			fmt.Println("Channel 2 has been closed")
			close(ch2)
			ch1 <- val + 1
			goto skipFor
			//break
		} else {
			fmt.Println("Even Number is:", val)
			ch1 <- val + 1
		}

	}
skipFor:
	fmt.Println("Skipped For loop showEvenNumbers")
	//ch1 <- val + 1
}
