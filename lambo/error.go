package lambgo

import (
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func NewHTTPError(msg string) (events.APIGatewayResponse, error) {
	return events.APIGatewayResponse{
		StatusCode: http.StatusInternalServerError,
	}, errors.New(msg)
}
