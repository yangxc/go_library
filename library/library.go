package main

import (
	"library/config"
	"log"
)

func main() {
	log.Println("Starting Library Server")
	log.Println("Initializing configuration")
	err := config.InitConfig("library", nil)
	if err != nil {
		log.Fatalf("Failed to read configuration: %v\n", err)
	}
	log.Println("Library Server Stopped")
}
