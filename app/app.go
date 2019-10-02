package app

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chiips/password/english/generator"
	"github.com/chiips/password/francais/generateur"
	"github.com/chiips/password/italiano/generatore"
	"github.com/urfave/cli"
)

//App is our CLI struct on which all set up functions hang
type App struct {
	*cli.App
}

//package-level variable for language flag
var language string

//package-level variable for easily testing case of language unavailable
var unavailable bool

//NewApp returns new app from urfave/cli
func NewApp() *App {

	app := cli.NewApp()

	return &App{app}

}

//SetInfo establishes info for the app
func (app *App) SetInfo() {
	app.Name = "Password Generator"
	app.Usage = "Generate cryptographically secure passwords."
	app.Author = "chiips"
	app.Version = "1.0.0"
}

//SetFlags establish CLI flags
func (app *App) SetFlags() {

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "language, lang, l",
			Value:       "english",
			Usage:       "language for generator instructions (options: \"francais\", \"italiano\")",
			Destination: &language,
		},
	}

}

//SetCommands establish CLI commands
func (app *App) SetCommands() {
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"gen", "g"},
			Usage:   "Generates a password",
			Flags:   app.Flags,
			Action: func(c *cli.Context) error {

				//set reader to ask questions and handle responses
				reader := bufio.NewReader(os.Stdin)

				if language == "english" {

					fmt.Print("Welcome to your Password Generator.\n\nHow would you like your password?\nPlease type 'y' or 'n' to give your responses.\n\n")

					password, err := generator.Generate(reader)
					if err != nil {
						return err
					}

					fmt.Println("\nBehold your password:", password)
					return nil
				} else if language == "francais" {

					fmt.Print("Bienvenue à votre Générateur de Mots de Passe.\n\nComment voulez-vous votre mot de passe?\nVeuillez taper 'o' ou 'n' pour répondre.\n\n")

					password, err := generateur.Generer(reader)
					if err != nil {
						return err
					}

					fmt.Println("\nVoici votre mot de passe:", password)
					return nil
				} else if language == "italiano" {

					fmt.Print("Benvenuto al suo Generatore di Password.\n\nCome vorrebbe la sua password?\nSi prega di digitare 's' o 'n' per rispondere.\n\n")

					password, err := generatore.Generare(reader)
					if err != nil {
						return err
					}

					fmt.Println("\nEcco la sua password:", password)
					return nil
				}

				//set unavailable to true for easy testing
				unavailable = true
				//respond to user
				fmt.Print("Requested language uavailable. Please choose from available languages or exclude the language flag for English.\n\n")
				err := cli.ShowCommandHelp(c, "generate")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}
