package main

import (
	"log"

	"github.com/qlfzn/goffee-club/cmd/api"
)

func main() {
	app := api.Application{
		Addr: ":8080",
	}

	r := app.Mount()

	err := app.Run(r)
	if err != nil {
		log.Fatalf("error running server: %s", err)
	}
}