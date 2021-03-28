package main

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		err     error
		request Request
		expect  Response
	}{
		{
			// Test that the handler responds with the correct response
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
