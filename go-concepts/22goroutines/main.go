package main

import (
	"fmt"
	"net/http"
	"sync"
)

// go routines helps to achive parallalism
// it is alternative of threads(managed by OS, 1MB fixed size), goroutines(managed by GO runtime,2kb flexible size)
var signals = []string{"test"}
var wg sync.WaitGroup

// mutex - mutual exclusion lock enables to share same memory by multiple go routines
var mut sync.Mutex

func main() {
	// go greeter("Hello")
	// greeter("world")

	websites := []string{
		"http://lco.dev",
		"http://google.com",
		"http://fb.com",
		"http://github.com",
		"http://go.dev",
		"http://youtube.com",
	}
	for _, web := range websites {
		go getStatusCode(web) // go routine created
		wg.Add(1)             // adds to the waitgroup (kind of registration of go routine)
	}

	wg.Wait() // this function will execute only at the end of the method, when all the go routines will get executed.
}

//	func greeter(s string) {
//		for i := 0; i < 6; i++ {
//			time.Sleep(3 * time.Millisecond)
//			fmt.Println(s)
//		}
//	}
func getStatusCode(endpoint string) {
	defer wg.Done() // will be called at the end of the function saying that, this go routine is executed

	res, err := http.Get(endpoint)
	if err != nil {
		panic(err)
	}
	mut.Lock() // locking the memory so that only one at a time should use the memory (Critical section)
	signals = append(signals, endpoint)
	mut.Unlock() // unlocking the memory
	fmt.Println(res.StatusCode, endpoint)
}
