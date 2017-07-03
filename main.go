package main

import (
	"log"

	"github.com/kapmahc/h2o/lib/app"
	_ "github.com/lib/pq"
)

func main() {
	if err := app.Main(); err != nil {
		log.Fatal(err)
	}
}
