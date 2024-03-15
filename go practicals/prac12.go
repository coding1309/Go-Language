package main

import (
	"flag"
	"fmt"
)

// Define a custom configuration struct
type Config struct {
	ServerAddress string
	Port          int
	LogLevel      string
}

// Custom flag type for LogLevel
type LogLevelFlag struct {
	value string
}

// Implement the flag.Value interface methods
func (l *LogLevelFlag) String() string {
	return l.value
}

func (l *LogLevelFlag) Set(value string) error {
	validLevels := map[string]bool{"debug": true, "info": true, "warn": true, "error": true}
	if _, ok := validLevels[value]; !ok {
		return fmt.Errorf("invalid log level: %s", value)
	}
	l.value = value
	return nil
}

func main() {
	// Define flags with default values
	var serverAddress = flag.String("server", "localhost", "Server address")
	var port = flag.Int("port", 8080, "Server port")
	var logLevel = &LogLevelFlag{"info"} // Use pointer for custom flag
	flag.Var(logLevel, "loglevel", "Log level (debug, info, warn, error)")

	// Parse command-line arguments
	flag.Parse()

	// Create the configuration object
	config := Config{
		ServerAddress: *serverAddress,
		Port:          *port,
		LogLevel:      logLevel.value,
	}

	// Print the configuration
	fmt.Printf("Server Address: %s\n", config.ServerAddress)
	fmt.Printf("Port:            %d\n", config.Port)
	fmt.Printf("Log Level:        %s\n", config.LogLevel)
}