package main

import (
	"log"
	"os"

	"github.com/chiips/password/app"
)

func main() {

	app := app.NewApp()

	app.SetInfo()
	app.SetFlags()
	app.SetCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	
}
