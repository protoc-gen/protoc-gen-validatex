package main

import (
	"log"
)

func main() {
	// Create a SignInRequest instance
	req := &SignInRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Validate the SignInRequest instance
	if err := req.Validate(); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	log.Println("validation passed")
}
