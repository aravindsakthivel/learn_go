package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}

func crxWithCancel() {
	crx := context.Background()
	crx, cancel := context.WithCancel(crx)

	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		cancel()
	}()

	sleepAndTalk(crx, 5*time.Second, "Hello")
}

func crxWithTimeout() {
	crx := context.Background()

	crx, cancel := context.WithTimeout(crx, 2*time.Second)

	defer cancel()

	sleepAndTalk(crx, 3*time.Second, "Hello")

}

func crxWithDeadline() {
	crx := context.Background()

	crx, cancel := context.WithDeadline(crx, time.Now().Add(2*time.Second))

	defer cancel()

	sleepAndTalk(crx, 3*time.Second, "Hello")
}

func processData(ctx context.Context) {
	// Normally, you would retrieve values from the context or handle cancellation
	// Since this is a placeholder, we're just simulating a process

	fmt.Println("Processing data with placeholder context")

	// Simulate checking for a value that would be added to a real context
	value := ctx.Value("key")

	if value != nil {
		fmt.Printf("Found value in context: %v\n", value)
	} else {
		fmt.Println("No value found in context")
	}
}

func main() {
	// crxWithCancel()
	// crxWithTimeout()
	crxWithDeadline()

	// Create a new context
	ctx := context.TODO()

	processData(ctx)
}
