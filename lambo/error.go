package lambo

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func NewHTTPError(msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,
	}, errors.New(msg)
}
