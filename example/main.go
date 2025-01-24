package main

import (
	"context"
	"github.com/protoc-gen/protoc-gen-validatex/pkg/validatex"
	"log"
)

func main() {
	// Create a SignInRequest instance
	req := &SignInRequest{
		Email:    "test.com",
		Password: "password123",
	}

	ctx := context.WithValue(context.Background(), validatex.KeyXLang, "zh")

	// Validate the SignInRequest instance
	if err := req.Validate(ctx); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	log.Println("validation passed")
}
