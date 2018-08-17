package github

import (
	"github.com/ntrv/lambo/lambo"
	gh "gopkg.in/go-playground/webhooks.v3/github"
)

type Hook struct {
	eventName  gh.Event
	eventFuncs map[gh.Event]lambo.HandleProcessFunc
	postFuncs  map[gh.Event]lambo.PostProcessFunc
}

type Event struct {
	EventName gh.Event
	HandleProcess lambo.HandleProcessFunc
	PostProcess lambo.PostProcessFunc
}

func NewHook() Hook {
	return Hook{
		eventFuncs: map[gh.Event]lambo.HandleProcessFunc{},
		postFuncs:  map[gh.Event]lambo.PostProcessFunc{},
	}
}

func (h *Hook) RegisterEvents(events ...Event) {
	for _, e := range events {
		h.eventFuncs[e.EventName] = e.HandleProcess
		h.postFuncs[e.EventName] = e.PostProcess
	}
}
