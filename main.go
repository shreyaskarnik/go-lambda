package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
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
