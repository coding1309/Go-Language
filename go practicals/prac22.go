package main

import (
	"fmt"
	"sync"
)

// Config represents the configuration object
type Config struct {
	// Add fields for configuration parameters
	SomeParameter string
	AnotherParam  int
}

// ConfigLoader loads the configuration
type ConfigLoader struct {
	once     sync.Once
	config   *Config
	loadFunc func() *Config // Function to load the configuration
}

// Load loads the configuration
func (cl *ConfigLoader) Load() *Config {
	cl.once.Do(func() {
		// Replace this with your configuration loading logic
		cl.config = &Config{
			SomeParameter: "default value",
			AnotherParam:  42,
		}
		fmt.Println("Configuration loaded")
	})
	return cl.config
}

func main() {
	// Create a ConfigLoader instance
	loader := &ConfigLoader{
		loadFunc: loadConfiguration,
	}

	// Access the configuration concurrently
	var wg sync.WaitGroup
	wg.Add(2)

	// Two goroutines accessing the configuration
	go func() {
		defer wg.Done()
		config := loader.Load()
		fmt.Println("Goroutine 1:", config)
	}()

	go func() {
		defer wg.Done()
		config := loader.Load()
		fmt.Println("Goroutine 2:", config)
	}()

	wg.Wait()
}

// Example function to load configuration
func loadConfiguration() *Config {
	// Simulate loading configuration from external source
	return &Config{
		SomeParameter: "loaded value",
		AnotherParam:  100,
	}
}
