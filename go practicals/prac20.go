package main

import (
	"fmt"
	"runtime"
	"time"
)

var counter int

func increment(id int) {
	counter++
	fmt.Println("Goroutine", id, "counter:", counter)
}

func main() {
	runtime.GOMAXPROCS(2) // Allow up to 2 CPUs for concurrency

	for i := 0; i < 10; i++ {
		go increment(i)
	}

	time.Sleep(time.Second) // Wait for goroutines to finish (may or may not be enough)
	fmt.Println("Final counter:", counter)
}
