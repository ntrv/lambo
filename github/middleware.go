package github

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/ntrv/lambo/lambo"
	"github.com/prometheus/prometheus/config"
)

// MiddlewareVerify ... Verify X-Hub-Signature and secret
func MiddlewareVerify(secret string) lambo.MiddlewareFunc {
	return func(next lambo.HandlerFunc) lambo.HandlerFunc {
		return func(ctx context.Context, req events.APIGatewayProxyRequest) (
			events.APIGatewayProxyResponse, error) {
			// Pick up GitHub Signature from header
			// And remove "sha1=" from X-Hub-Signature
			// See https://developer.github.com/webhooks/#delivery-headers
			signature := strings.TrimLeft(
				req.Headers["X-Hub-Signature"],
				"sha1=",
			)
			if len(signature) == 0 {
				return lambo.NewHTTPError("Missing X-Hub-Signature required for HMAC verification")
			}

			// Calculate hmac from HTTP body and secret key
			mac := hmac.New(sha1.New, []byte(secret))
			payload, err := ioutil.ReadAll(req.Body)
			if err != nil || len(payload) == 0 {
				return lambo.NewHTTPError("Issue reading Payload")
			}
			mac.Write(payload)
			expectedMac := hex.EncodeToString(mac.Sum(nil))

			// Compare whether signature matches calculated value
			if !hmac.Equal([]byte(signature), []byte(expectedMac)) {
				return lambo.NewHTTPError("HMAC verification failed")
			}
			return next(ctx, req)
		}
	}
}

func MiddlewareCheckMethod(next lambo.HandlerFunc) lambo.HandlerFunc {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (
		events.APIGatewayProxyResponse, error) {
		if req.HTTPMethod != "POST" {
			return lambo.NewHTTPError("")
		}
		return next(ctx, req)
	}
}
