package github

import (
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

type Hook struct {
	eventName  gh.Event
	eventFuncs map[gh.Event]lambo.HandleProcessFunc
}

type Event struct {
	EventName gh.Event
	HandleProcess lambo.HandleProcessFunc
}

func NewHook() Hook {
	return Hook{
		eventFuncs: map[gh.Event]lambo.HandleProcessFunc{},
	}
}

func (h *Hook) RegisterEvents(events ...Event) {
	for _, e := range events {
		h.eventFuncs[e.EventName] = e.HandleProcess
	}
}
