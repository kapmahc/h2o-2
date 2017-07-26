package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kapmahc/axe"
	_ "github.com/kapmahc/axe/base"
)

func main() {
	if err := axe.Main(); err != nil {
		log.Panic(err)
	}
}
