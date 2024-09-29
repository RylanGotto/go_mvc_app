package main

import (
	"app/app1/handlers"

	"github.com/rylangotto/tonic"
)

type application struct {
	App      *tonic.Tonic
	Handlers *handlers.Handlers
}

func main() {
	t := initApplication()
	t.App.ListenAndServer()
}
