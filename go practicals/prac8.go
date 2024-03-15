package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var num int
	for {
		fmt.Print("Enter a number to calculate its factorial (or enter -1 to exit): ")
		fmt.Scanln(&num)
		if num == -1 {
			break
		}
		fmt.Printf("Factorial of %d is %d\n", num, factorial(num))
	}
}
