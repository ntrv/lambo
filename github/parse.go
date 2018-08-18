package github

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func (h *Hook) ParsePayloadHandler(ctx context.Context, req events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {

	event := req.Headers["X-GitHub-Event"]
	if len(event) == 0 {
		return lambo.NewHTTPError(
			http.StatusBadRequest,
			"Missing X-GitHub-Event Header",
		)
	}

	h.eventName = gh.Event(event)
	fn, ok := h.eventFuncs[h.eventName]
	if !ok {
		return lambo.NewHTTPError(
			http.StatusNotImplemented,
			fmt.Sprintf(
				"Webhook Event %s not registered, it is recommended to setup only events in github that will be registered in the webhook to avoid unnecessary traffic and reduce potential attack vectors.",
				event,
			),
		)
	}
	return h.runProcessContext(ctx, fn, req)
}
