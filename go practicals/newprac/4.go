package main

import (
	"fmt"
)

func printTasks(tasks []string) {
	fmt.Println("Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func addTask(tasks *[]string, task string) {
	*tasks = append(*tasks, task)
}

func removeTask(tasks *[]string, index int) {
	if index < 1 || index > len(*tasks) {
		fmt.Println("Invalid task index")
		return
	}
	*tasks = append((*tasks)[:index-1], (*tasks)[index:]...)
}

func main() {
	var tasks []string
	for {
		fmt.Println("1. Add task")
		fmt.Println("2. Remove task")
		fmt.Println("3. Display tasks")
		fmt.Println("4. Exit")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task: ")
			var task string
			fmt.Scanln(&task)
			addTask(&tasks, task)
			fmt.Println("Task added")
		case 2:
			fmt.Print("Enter task index to remove: ")
			var index int
			fmt.Scanln(&index)
			removeTask(&tasks, index)
			fmt.Println("Task removed")
		case 3:
			printTasks(tasks)
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
