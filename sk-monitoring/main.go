package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/thenguyenit/sk-monitoring/cron"
)

func main() {
	fmt.Println("Start to run a goroutine")

	//Read the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Verify the log
	cron.Monitor()
}
