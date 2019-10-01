package main

import (
	"log"
	"os"

	"github.com/chiips/password/app"
)

func main() {

	cli := app.NewApp()

	cli.SetInfo()
	cli.SetFlags()
	cli.SetCommands()

	err := cli.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
