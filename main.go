package main

import (
	"go-graphql/app"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	app := app.NewApp()
	app.Run()
}
