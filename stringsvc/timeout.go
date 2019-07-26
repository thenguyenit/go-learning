package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	log.Println("Starting...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Case 1")

	select {
	case <-time.After(10 * time.Second):
		log.Println("overslept")
	// fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err()) // prints "context deadline exceeded"
	}
}
