package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Request is the struct for Request Object
type Request struct {
	ID    float64 `json:"id"`
	Value string  `json:"value"`
}

// Response is the struct Response Object
type Response struct {
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Handler function for lambda runtime
func Handler(request Request) (Response, error) {
	return Response{
		Message: fmt.Sprintf("Processed request ID %f", request.ID),
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
