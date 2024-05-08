package main

import (
	"fmt"
)

func main() {
	var numScores int

	fmt.Println("Enter the number of exam scores:")
	fmt.Scanln(&numScores)

	if numScores <= 0 {
		fmt.Println("Invalid number of scores. Please enter a positive integer.")
		return
	}

	scores := make([]float64, numScores)

	fmt.Println("Enter exam scores:")
	for i := 0; i < numScores; i++ {
		fmt.Printf("Score %d: ", i+1)
		fmt.Scanln(&scores[i])
	}

	var total float64
	for _, score := range scores {
		total += score
	}

	average := total / float64(numScores)
	fmt.Printf("The average score is %.2f\n", average)
}
