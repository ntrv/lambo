package github

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/ntrv/lambo/lambo"
)

// MiddlewareVerify ... Verify X-Hub-Signature and secret
func MiddlewareVerify(secret string) lambo.MiddlewareFunc {
	return func(next lambo.HandlerFunc) lambo.HandlerFunc {
		return func(ctx context.Context, req events.APIGatewayProxyRequest) (
			events.APIGatewayProxyResponse, error) {

			// Use AWS X-Ray
			ctx, seg := xray.BeginSegment(ctx, "VerifySignature")
			defer seg.Close(nil)

			// Pick up GitHub Signature from header
			// And remove "sha1=" from X-Hub-Signature
			// See https://developer.github.com/webhooks/#delivery-headers
			signature := strings.TrimLeft(
				req.Headers["X-Hub-Signature"],
				"sha1=",
			)
			if len(signature) == 0 {
				return lambo.NewHTTPError(
					http.StatusBadRequest,
					"Missing X-Hub-Signature required for HMAC verification",
				)
			}

			// Calculate hmac from HTTP body and secret key
			if len(req.Body) == 0 {
				return lambo.NewHTTPError(
					http.StatusNotAcceptable,
					"Missing Body",
				)
			}
			mac := hmac.New(sha1.New, []byte(secret))
			mac.Write([]byte(req.Body))
			expectedMac := hex.EncodeToString(mac.Sum(nil))

			// Compare whether signature matches calculated value
			if !hmac.Equal([]byte(signature), []byte(expectedMac)) {
				return lambo.NewHTTPError(
					http.StatusForbidden,
					"HMAC verification failed",
				)
			}
			return next(ctx, req)
		}
	}
}

func MiddlewareCheckMethod(next lambo.HandlerFunc) lambo.HandlerFunc {
	return func(ctx context.Context, req events.APIGatewayProxyRequest) (
		events.APIGatewayProxyResponse, error) {
		ctx, seg := xray.BeginSegment(ctx, "CheckMethod")
		defer seg.Close(nil)

		if req.HTTPMethod != http.MethodPost {
			return lambo.NewHTTPError(
				http.StatusMethodNotAllowed,
				fmt.Sprintf(
					"Attempt made using following method is not allowed: %s",
					req.HTTPMethod,
				),
			)
		}
		return next(ctx, req)
	}
}
