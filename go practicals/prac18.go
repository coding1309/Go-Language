package main

import (
    "fmt"
    "sync"
)

type ChatRoom struct {
    name     string
    messages chan string
}

func NewChatRoom(name string) *ChatRoom {
    return &ChatRoom{
        name:     name,
        messages: make(chan string),
    }
}

func (cr *ChatRoom) Join(user string) {
    message := fmt.Sprintf("%s joined the chat room %s", user, cr.name)
    cr.messages <- message
}

func (cr *ChatRoom) Leave(user string) {
    message := fmt.Sprintf("%s left the chat room %s", user, cr.name)
    cr.messages <- message
}

func (cr *ChatRoom) Broadcast() {
    for msg := range cr.messages {
        fmt.Printf("[%s] %s\n", cr.name, msg)
    }
}

func main() {
    // Create chat rooms
    room1 := NewChatRoom("Room 1")
    room2 := NewChatRoom("Room 2")

    // Wait group to wait for all goroutines to finish
    var wg sync.WaitGroup
    wg.Add(2)

    // Goroutines to handle broadcasting in each chat room
    go func() {
        defer wg.Done()
        room1.Broadcast()
    }()

    go func() {
        defer wg.Done()
        room2.Broadcast()
    }()

    // Simulate users joining and leaving chat rooms
    room1.Join("User 1")
    room2.Join("User 2")
    room1.Join("User 3")
    room2.Leave("User 2")
    room1.Leave("User 1")

    // Close channels to signal that no more messages will be sent
    close(room1.messages)
    close(room2.messages)

    // Wait for goroutines to finish
    wg.Wait()
}
