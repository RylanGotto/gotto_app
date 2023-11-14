package main

import (
	"app/handlers"

	"github.com/RylanGotto/gotto"
)

type application struct {
	App      *gotto.Gotto
	Handlers *handlers.Handlers
}

func main() {
	g := initApplication()

	g.App.ListenAndServe()
}
