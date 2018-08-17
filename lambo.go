package main

import "github.com/ntrv/lambo/lambo"

func main() {
	l := lambo.New()
	l.Use(lambo.MiddlewareExample)
	l.Run(lambo.HandleEcho)
}
