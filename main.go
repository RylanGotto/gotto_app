package main

import "github.com/RylanGotto/gotto"

type application struct {
	App *gotto.Gotto
}

func main() {
	g := initApplication()
	g.App.ListenAndServe()
}
