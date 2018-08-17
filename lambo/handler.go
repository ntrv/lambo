package lambo

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// HandleEcho ... Just echo handler
func HandleEcho(ctx context.Context, req events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: req.Body, StatusCode: http.StatusOK}, nil
}
