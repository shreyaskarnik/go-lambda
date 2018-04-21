package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request Request
		expect  Response
		err     error
	}{
		{
			// Test that the handler responds with the correct response
			// when a valid name is provided in the HTTP body
			request: Request{
				ID:    math.Pi,
				Value: "pi",
			},
			expect: Response{
				Message: fmt.Sprintf("Processed request ID %f", math.Pi),
				Ok:      true,
			},
			err: nil,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expect, response)
	}
}
