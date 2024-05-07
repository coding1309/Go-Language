package main

import (
	"fmt"
)

func performOperation(operation rune) float64 {
	var num1, num2 float64
	fmt.Println("Enter two numbers and an operation (+ or * or - or / ) to perform:")
	fmt.Print("First number: ")
	fmt.Scanf("%f\n", &num1)
	fmt.Print("Second number: ")
	fmt.Scanf("%f\n", &num2)

	switch operation {
	case '+':
		return num1 + num2
	case '*':
		return num1 * num2
	case '-':
		return num1 - num2
	case '/':
		if num2 == 0 {
			fmt.Println("Error: Division by zero")
			return 0
		}
		return num1 / num2
	case 'N':
		fmt.Println("Exiting...")
		return 0
	default:
		fmt.Println("Invalid operation. Supported operations are +, -, *, /, and N for Exit.")
		return 0
	}
}

func main() {
	var operation rune
	for {
		fmt.Println("+ For Additon.")
		fmt.Println("- For Subtraction.")
		fmt.Println("/ For Division.")
		fmt.Println("* For Multiplication.")
		fmt.Println("N for Exit")
		fmt.Scanf("%c\n", &operation)
		result := performOperation(operation)
		if result != 0 {
			fmt.Printf("Result: %.2f\n", result)
		}
	}
}
