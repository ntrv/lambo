package lambo

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type HandleProcessFunc func(context.Context, interface{}, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
type PostProcessFunc func(string, events.APIGatewayProxyRequest) error
