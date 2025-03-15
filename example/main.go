package main

import (
	"context"
	"log"
	"time"

	"github.com/protoc-gen/protoc-gen-validatex/pkg/validatex"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	req := &Playground{
		Email:       "a@test.com",
		Username:    "test",
		Uuid:        "123e4567-e89b-12d3-a456-426614174000",
		Score:       100,
		Age:         10,
		Temperature: 30,
		CreatedAt: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		},
		UpdatedAt: &timestamppb.Timestamp{
			Seconds: time.Now().Unix(),
			Nanos:   0,
		},
		ExpiresAt: &timestamppb.Timestamp{
			Seconds: time.Now().Add(time.Hour).Unix(),
			Nanos:   0,
		},
		EventTime: &timestamppb.Timestamp{
			Seconds: 1672531200,
			Nanos:   0,
		},
		Tags:          []string{"tag1", "tag2"},
		Scores:        []int32{100, 200, 300},
		TermsAccepted: true,
		IsAdult:       true,
	}

	ctx := context.WithValue(context.Background(), validatex.KeyXLang, "zh")

	if err := req.Validate(ctx); err != nil {
		log.Fatalf("validation failed: %v", err)
	}

	log.Println("validation passed")
}
