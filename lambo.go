package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/ntrv/lambo/github"
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	xray.Configure(xray.Config{LogLevel: "warn"})

	hook := github.NewHook()
	hook.RegisterEvents(github.Event{
		EventName:     gh.PushEvent,
		HandleProcess: lambo.HandlePushSample,
	})

	l := lambo.New()
	l.Use(lambo.MiddlewareExample)
	lambda.Start(l.Then(hook.ParsePayloadHandler))
}
