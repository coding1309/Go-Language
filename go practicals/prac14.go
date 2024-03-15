package main

import (
	"fmt"
	"sort"
)

// Person is a custom type to represent a person with a name, age, and city.
type Person struct {
	Name string
	Age  int
	City string
}

// ByName implements the sort.Interface for sorting a slice of Person structs by name.
type ByName []Person

// Len implements the sort.Interface.
func (b ByName) Len() int { return len(b) }

// Swap implements the sort.Interface.
func (b ByName) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements the sort.Interface.
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }

// ByAge implements the sort.Interface for sorting a slice of Person structs by age.
type ByAge []Person

// Len implements the sort.Interface.
func (b ByAge) Len() int { return len(b) }

// Swap implements the sort.Interface.
func (b ByAge) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements the sort.Interface.
func (b ByAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

// ByCity implements the sort.Interface for sorting a slice of Person structs by city.
type ByCity []Person

// Len implements the sort.Interface.
func (b ByCity) Len() int { return len(b) }

// Swap implements the sort.Interface.
func (b ByCity) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

// Less implements the sort.Interface.
func (b ByCity) Less(i, j int) bool { return b[i].City < b[j].City }

func main() {
	people := []Person{
		{"Neel", 22, "Morbi"},
		{"Ravi", 28, "Tankara"},
		{"Kishan", 35, "Rajkot"},
		{"Biren", 40, "Mehsana"},
	}

	fmt.Println("Before sorting:")
	fmt.Println(people, "\n")

	// Sort by name
	sort.Sort(ByName(people))
	fmt.Println("After sorting by name:")
	fmt.Println(people, "\n")

	// Sort by age
	sort.Sort(ByAge(people))
	fmt.Println("After sorting by age:")
	fmt.Println(people, "\n")

	// Sort by city
	sort.Sort(ByCity(people))
	fmt.Println("After sorting by city:")
	fmt.Println(people, "\n")
}
