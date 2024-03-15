package main

import (
	"fmt"
)

// Container defines a generic container interface
type Container interface {
	Add(interface{})
	Get(int) interface{}
	Len() int
}

// MyContainer implements the Container interface
type MyContainer struct {
	data []interface{}
}

// NewMyContainer creates a new instance of MyContainer
func NewMyContainer() *MyContainer {
	return &MyContainer{
		data: make([]interface{}, 0),
	}
}

// Add adds a value to the container
func (c *MyContainer) Add(value interface{}) {
	c.data = append(c.data, value)
}

// Get retrieves a value from the container by index
func (c *MyContainer) Get(index int) interface{} {
	return c.data[index]
}

// Len returns the number of elements in the container
func (c *MyContainer) Len() int {
	return len(c.data)
}

func main() {
	container := NewMyContainer()

	// Adding values of different types
	container.Add(42)
	container.Add("Hello")
	container.Add(3.14)

	// Retrieving and printing values
	for i := 0; i < container.Len(); i++ {
		fmt.Println(container.Get(i))
	}
}
