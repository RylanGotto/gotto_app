package main

import (
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

	app := &application{
		App: got,
	}

	return app
}
