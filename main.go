package main

import (
	"log"

	"github.com/jffcm/aluraflix-backend/config"
	"github.com/jffcm/aluraflix-backend/router"
)

func main() {
	// Initialize Configs
	err := config.Init()
	if err != nil {
		log.Printf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
