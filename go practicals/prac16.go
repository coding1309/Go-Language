package main

import (
	"errors"
	"fmt"
	"os"
)

// Define custom error types
var (
	ErrFileNotFound     = errors.New("file not found")
	ErrPermissionDenied = errors.New("permission denied")
	ErrUnknownError     = errors.New("unknown error")
)

func readFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrFileNotFound
		} else if os.IsPermission(err) {
			return nil, ErrPermissionDenied
		}
		return nil, fmt.Errorf("read error: %w", err) // Wrap original error
	}
	return data, nil
}

func writeFile(path string, data []byte) error {
	err := os.WriteFile(path, data, 0644) // Set appropriate permissions
	if err != nil {
		if os.IsPermission(err) {
			return ErrPermissionDenied
		}
		return fmt.Errorf("write error: %w", err) // Wrap original error
	}
	return nil
}

func main() {
	// Read a file
	data, err := readFile("data.txt")
	if err != nil {
		switch err {
		case ErrFileNotFound:
			fmt.Println("File not found:", err)
		case ErrPermissionDenied:
			fmt.Println("Permission denied:", err)
		default:
			fmt.Println("Error reading file:", err)
		}
		return
	}
	fmt.Println(string(data))

	// Write to a file
	err = writeFile("output.txt", data)
	if err != nil {
		if err == ErrPermissionDenied {
			fmt.Println("Permission denied writing to file:", err)
		} else {
			fmt.Println("Error writing to file:", err)
		}
		return
	}
	fmt.Println("File written successfully")
}
