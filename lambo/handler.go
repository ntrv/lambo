package lambo

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/xray"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

type HandleProcessFunc func(context.Context, interface{}, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// HandleEcho ... Just echo handler
func HandleEcho(ctx context.Context, req events.APIGatewayProxyRequest) (
	events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{Body: req.Body, StatusCode: http.StatusOK}, nil
}

func HandlePushSample(
	ctx context.Context,
	payload interface{},
	req events.APIGatewayProxyRequest,
) (events.APIGatewayProxyResponse, error) {

	// Use AWS X-Ray
	ctx, seg := xray.BeginSegment(ctx, "HandlerPush")
	defer seg.Close(nil)

	pl, ok := payload.(gh.PushPayload)
	if !ok {
		return NewHTTPError(
			http.StatusInternalServerError,
			"Failed to parse PushPayload",
		)
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       pl.Commits[0].Message,
	}, nil
}
