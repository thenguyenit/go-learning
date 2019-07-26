package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Started...")
	now := time.Now()

	ctx := context.Background()

	operationTimeout(ctx)

	elapsed := time.Since(now)
	fmt.Printf("Ended...After %s \n", elapsed)
}

func operationTimeout(ctx context.Context) {
	fmt.Println("operationTimeout Started...")

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	c := make(chan string)

	go slowOperation(c)

	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-c:
		fmt.Println(c)
	}

	fmt.Println("operationTimeout Ended...")
}

func slowOperation(c chan string) {
	fmt.Println("slowOperation Started...")
	time.Sleep(3 * time.Second)
	c <- "Tada..."
	fmt.Println("slowOperation Ended...")
}
