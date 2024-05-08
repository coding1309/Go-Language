package main

import (
	"fmt"
)

func main() {
	var temperature float64
	var unit string

	fmt.Println("Enter temperature value:")
	fmt.Scanln(&temperature)

	fmt.Println("Enter unit (C for Celsius, F for Fahrenheit):")
	fmt.Scanln(&unit)

	var convertedTemp float64
	switch unit {
	case "C":
		convertedTemp = (temperature * 1.8) + 32
		fmt.Printf("%.2f degrees Celsius is equal to %.2f degrees Fahrenheit\n", temperature, convertedTemp)
	case "F":
		convertedTemp = (temperature - 32) * 5 / 9
		fmt.Printf("%.2f degrees Fahrenheit is equal to %.2f degrees Celsius\n", temperature, convertedTemp)
	default:
		fmt.Println("Invalid unit. Please enter 'C' or 'F'.")
	}
}
