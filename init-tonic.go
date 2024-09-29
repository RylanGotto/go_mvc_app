package main

import (
	"app/app1/handlers"
	"log"
	"os"

	"github.com/rylangotto/tonic"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init tonic
	ton := &tonic.Tonic{}
	err = ton.New(path)
	if err != nil {
		log.Fatal(err)
	}

	ton.AppName = "app"

	myHandlers := &handlers.Handlers{
		App: ton,
	}

	app := &application{
		App:      ton,
		Handlers: myHandlers,
	}

	app.App.Routes = app.routes()

	return app

}
