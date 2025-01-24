package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignInRequest_ValidateEmail(t *testing.T) {
	// Define test cases
	tests := []struct {
		name        string
		req         *SignInRequest
		expectError bool
	}{
		{
			name: "Valid input",
			req: &SignInRequest{
				Email:    "test@example.com",
				Password: "password123",
			},
			expectError: false,
		},
		{
			name: "Empty email",
			req: &SignInRequest{
				Email:    "",
				Password: "password123",
			},
			expectError: true,
		},
		{
			name: "Empty password",
			req: &SignInRequest{
				Email:    "test@example.com",
				Password: "",
			},
			expectError: true,
		},
		{
			name: "Empty email and password",
			req: &SignInRequest{
				Email:    "",
				Password: "",
			},
			expectError: true,
		},
		{
			name: "Missing '@' symbol",
			req: &SignInRequest{
				Email:    "testexample.com",
				Password: "password123",
			},
			expectError: true,
		},
		{
			name: "Path too long",
			req: &SignInRequest{
				Email:    "ab@0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789.com",
				Password: "password123",
			},
			expectError: true,
		},
		{
			name: "Part before '@' too long",
			req: &SignInRequest{
				Email:    "01234567890123456789012345678901234567890123456789012345678901234@domain.com",
				Password: "password123",
			},
			expectError: true,
		},
		{
			name: "Part after '@' too long",
			req: &SignInRequest{
				Email:    "a@01234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890.com",
				Password: "password123",
			},
			expectError: true,
		},
		{
			name: "No domain after '@'",
			req: &SignInRequest{
				Email:    "test@",
				Password: "password123",
			},
			expectError: true,
		},
		{
			name: "Invalid email format (missing dot)",
			req: &SignInRequest{
				Email:    "test@com",
				Password: "password123",
			},
			expectError: true,
		},
	}

	ctx := context.Background()

	// Iterate over each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the validation
			err := tt.req.Validate(ctx)

			// Check if the error is expected
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
