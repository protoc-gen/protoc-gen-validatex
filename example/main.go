package main

import (
	"context"
	"github.com/protoc-gen/protoc-gen-validatex/pkg/validatex"
	"log"
)

func main() {
	req := &PlaygroundRequest{
		Email:    "a@test.com",
		Password: "password123",
		Id:       "123e4567-e89b-12d3-a456-426614174000",
	}

	ctx := context.WithValue(context.Background(), validatex.KeyXLang, "zh")

	if err := req.Validate(ctx); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	log.Println("validation passed")
}
