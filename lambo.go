package main

import (
	"github.com/ntrv/lambo/github"
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

func main() {
	hook := github.NewHook()
	hook.RegisterEvents(github.Event{
		EventName:     gh.PushEvent,
		HandleProcess: lambo.HandlePushSample,
	})

	l := lambo.New()
	l.Use(lambo.MiddlewareExample)
	// l.Run(lambo.HandleEcho)
	l.Run(hook.ParsePayloadHandler)
}
