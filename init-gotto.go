package main

import (
	"app/handlers"
	"log"
	"os"

	"github.com/RylanGotto/gotto"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init gotto
	got := &gotto.Gotto{}
	err = got.New(path)
	if err != nil {
		log.Fatal(err)
	}
	got.AppName = "app"

	routeHandlers := &handlers.Handlers{
		App: got,
	}

	app := &application{
		App:      got,
		Handlers: routeHandlers,
	}

	app.App.Routes = app.routes()

	return app
}
