package main

import (
    "context"
    "fmt"
    "math/big"
    "time"
)

// performCalculation performs a lengthy calculation that can be canceled on user demand
func performCalculation(ctx context.Context) *big.Int {
    result := new(big.Int)
    result.SetInt64(0)

    // Simulate a lengthy calculation
    for i := int64(0); i < 1000000000; i++ {
        select {
        case <-ctx.Done():
            fmt.Println("Calculation canceled")
            return nil
        default:
            // Perform some computation
            result.Add(result, big.NewInt(i))
        }
    }

    return result
}

func main() {
    // Create a context with cancel function
    ctx, cancel := context.WithCancel(context.Background())

    // Perform the lengthy calculation in a Goroutine
    go func() {
        result := performCalculation(ctx)
        if result != nil {
            fmt.Println("Calculation result:", result)
        } else {
            fmt.Println("Calculation was canceled.")
        }
    }()

    // Wait for user input to cancel the calculation
    fmt.Println("Press 'Enter' to cancel the calculation.")
    fmt.Scanln()

    // Call cancel function to cancel the calculation
    cancel()

    // Wait for the Goroutine to finish
    time.Sleep(time.Second)
    fmt.Println("Program ended.")
}

