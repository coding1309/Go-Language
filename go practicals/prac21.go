package main

import (
	"fmt"
	"sync"
	"time"
)

type Config struct {
	value string
}

var config Config
var rwMutex sync.RWMutex // RWMutex for read/write access

func readConfig() (string, error) {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return config.value, nil
}

func writeConfig(newValue string) error {
	rwMutex.Lock()
	defer rwMutex.Unlock()
	config.value = newValue
	return nil
}

func main() {
	// Initial configuration value
	config.value = "Initial Value"

	// Launch multiple reader goroutines
	for i := 0; i < 5; i++ {
		go func(id int) {
			value, err := readConfig()
			if err != nil {
				fmt.Println("Error reading config:", err)
				return
			}
			fmt.Println("Goroutine", id, "read:", value)
		}(i)
	}

	// Launch a writer goroutine after a delay
	go func() {
		time.Sleep(time.Second * 2)
		err := writeConfig("Updated Value")
		if err != nil {
			fmt.Println("Error writing config:", err)
			return
		}
		fmt.Println("Wrote new configuration value")
	}()

	// Wait for all goroutines to finish
	time.Sleep(time.Second * 5)
}
