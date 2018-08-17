package lambo

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandlerFunc ...
type HandlerFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

type Lambo struct {
	middlewares []MiddlewareFunc
}

// New ...
func New() Lambo {
	return Lambo{middlewares: []MiddlewareFunc{}}
}

// Run ... Execute
func (l Lambo) Run(handler HandlerFunc) {
	lambda.Start(l.Then(handler))
}
