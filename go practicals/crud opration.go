package main

import (
	"fmt"
)

type employee struct {
	Id   int
	Name string
	City string
}

var employees []employee

func main() {
	var n int
	for {
		fmt.Println("1. Insert Data : ")
		fmt.Println("2. Show Data : ")
		fmt.Println("3. Delete Data : ")
		fmt.Println("4. Exit : ")

		fmt.Println("Enter Your Choice : ")
		fmt.Scanln(&n)

		switch n {
		case 1:
			insert()

		case 2:
			show()

		case 3:
			delete()
		case 4:
			fmt.Println("Exiting...Thank You  :)_")
			return
		}
	}
}

func insert() {
	var e employee
	fmt.Println("Enter an Employee Detalis : ")
	fmt.Print("Employee Id : ")
	fmt.Scanln(&e.Id)
	fmt.Print("Employee Name : ")
	fmt.Scanln(&e.Name)
	fmt.Print("Employee City : ")
	fmt.Scanln(&e.City)

	employees = append(employees, e)
	fmt.Println("Employee Added Successfully ..")
	fmt.Println("To Show Employee data plz press 2 .")
}
func show() {
	var id int
	find := false
	fmt.Print("Enter Id : ")
	fmt.Scanln(&id)
	for _, e := range employees {
		if e.Id == id {
			fmt.Println("Employee ID : ", e.Id)
			fmt.Println("Employee Name : ", e.Name)
			fmt.Println("Employee City : ", e.City)

			find = true
			break

		}
	}
	if !find {
		fmt.Printf("Not Found")
	}
}
func delete() {
	var id int
	find := false
	fmt.Print("Enter Id : ")
	fmt.Scanln(&id)
	for i, e := range employees {
		if e.Id == id {
			employees = append(employees[:i], employees[i+1:]...)
			find = true
			fmt.Println("Employee Data Deleted Successfully")
			break
		}
	}
	if !find {
		fmt.Println("Not Found")
	}

}
