package main

import (
	"fmt"
	"math/big"
	"runtime"
	"sync"
)

type FactorialThread struct {
	Num       int
	Factorial *big.Int
	Mutex     sync.Mutex
}

func factorialThread(n int, wg *sync.WaitGroup, threadId int, results chan FactorialThread) {
	defer wg.Done()

	var result big.Int
	result.SetUint64(1)
	for i := int64(1); i <= int64(n); i++ {
		result.Mul(&result, big.NewInt(i))
	}

	thread := FactorialThread{
		Num:       n,
		Factorial: &result,
	}

	results <- thread
}

func main() {
	numbers := []int{0, 1, 5, 7, 10}
	results := make(chan FactorialThread, len(numbers))
	var wg sync.WaitGroup

	runtime.GOMAXPROCS(len(numbers))

	for _, num := range numbers {
		wg.Add(1)
		go factorialThread(num, &wg, num, results)
	}

	wg.Wait()

	close(results)

	for thread := range results {
		fmt.Printf("Factorial of %d: %v\n", thread.Num, thread.Factorial)
	}

	fmt.Println("Done")
}
