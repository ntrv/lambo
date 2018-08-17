package lambo

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// MiddlewareFunc ... Middleware format
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// Use ... Include middlewares
func (l *Lambo) Use(mws ...MiddlewareFunc) {
	l.middlewares = append(l.middlewares, mws...)
}

// then ... Make middleware nested and generate a handler to actually use
func (l Lambo) then(h HandlerFunc) HandlerFunc {
	for i := range l.middlewares {
		h = l.middlewares[len(l.middlewares)-1-i](h)
	}
	return h
}

// MiddlewareExample ... Sample middleware function
func MiddlewareExample(next HandlerFunc) HandlerFunc {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (
		events.APIGatewayProxyResponse, error) {
		fmt.Printf("Processing request data %s.\n", req.RequestContext.RequestID)
		fmt.Printf("Body size = %d.\n", len(req.Body))
		fmt.Println("Headers:")
		for k, v := range req.Headers {
			fmt.Printf("    %s: %s\n", k, v)
		}
		return next(ctx, req)
	}
}
