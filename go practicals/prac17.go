// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // Task represents a unit of work to be executed concurrently
// type Task struct {
// 	URL    string
// 	Result chan string
// 	Error  chan error
// }

// func fetchData(task Task) {
// 	resp, err := http.Get(task.URL)
// 	if err != nil {
// 		task.Error <- err
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Simulate processing response
// 	data := fmt.Sprintf("Fetched data from %s", task.URL)
// 	task.Result <- data
// }

// func main() {
// 	urls := []string{
// 		"https://www.example.com/1",
// 		"https://www.example.com/2",
// 		"https://www.example.com/3",
// 	}

// 	// Create channels to communicate results and errors
// 	results := make(chan string)
// 	errors := make(chan error)

// 	// Launch goroutines for each task
// 	for _, url := range urls {
// 		task := Task{URL: url, Result: results, Error: errors}
// 		go fetchData(task)
// 	}

// 	// Process results and errors concurrently
// 	for i := range urls {
// 		select {
// 		case result := <-results:
// 			fmt.Println(result, i)
// 		case err := <-errors:
// 			fmt.Println("Error fetching data:", err)
// 		}
// 	}

// 	fmt.Println("Finished processing all tasks")
// }

package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string, ch chan string) {
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	ch <- fmt.Sprintf("Fetched %s with status code %d", url, resp.StatusCode)
}

func main() {
	urls := []string{
		"https://www.blackbox.ai/",
		"https://gemini.google.com",
		"https://chat.openai.com",
	}

	ch := make(chan string)

	// Launch a Goroutine for each URL
	for _, url := range urls {
		go fetchData(url, ch)
	}

	// Wait for all Goroutines to finish
	for range urls {
		fmt.Println(<-ch)
	}
}
