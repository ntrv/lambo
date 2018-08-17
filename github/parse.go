package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func (h *Hook) ParsePayloadHandler(ctx context.Context, req events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {
	if req.HTTPMethod != "POST" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("")
	}

	event := req.Headers["X-GitHub-Event"]
	if len(event) == 0 {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("")
	}

	h.eventName = gh.Event(event)
	fn, ok := h.eventFuncs[h.eventName]
	if !ok {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, fmt.Errorf("")
	}
	return h.runProcessContext(ctx, fn, req)
}
