package lambo

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
)

func NewHTTPError(code int, msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: code,
	}, errors.New(msg)
}
