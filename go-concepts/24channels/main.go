package main

import (
	"fmt"
	"sync"
)

func main() {
	mych := make(chan int, 2) // 2 is buffer channel, meaning it can store 2 variables
	wg := &sync.WaitGroup{}

	// mych <- 5 // 5 is going into channel
	// fmt.Println(<-mych) // getting out value from channel

	wg.Add(2)
	//recieve ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-mych

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-mych)

		wg.Done()
	}(mych, wg)

	//send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych)
		wg.Done()
	}(mych, wg)
	wg.Wait()
}
