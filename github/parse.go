package github

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func (h *Hook) ParsePayloadHandler(ctx context.Context, req events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {

	event := req.Headers["X-GitHub-Event"]
	if len(event) == 0 {
		return lambo.NewHTTPError("")
	}

	h.eventName = gh.Event(event)
	fn, ok := h.eventFuncs[h.eventName]
	if !ok {
		return lambo.NewHTTPError("")
	}
	return h.runProcessContext(ctx, fn, req)
}
