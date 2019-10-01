package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/chiips/password/english/generator"
	"github.com/chiips/password/francais/generateur"
	"github.com/chiips/password/italiano/generatore"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

var password, language string

var err error

var reader = bufio.NewReader(os.Stdin)

func info() {
	app.Name = "Password Generator"
	app.Usage = "Generate cryptographically secure passwords."
	app.Author = "chiips"
	app.Version = "1.0.0"
}

func flags() {

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "language, lang, l",
			Value:       "english",
			Usage:       "Language for generator instructions (options: 'francais', 'italiano')",
			Destination: &language,
		},
	}

}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "generate",
			Aliases: []string{"gen", "g"},
			Usage:   "Generates a password",
			Flags:   app.Flags,
			Action: func(c *cli.Context) error {

				if language == "francais" {
					fmt.Print("Bienvenue à votre Générateur de Mots de Passe.\n\nComment voudriez-vous votre mot de passe?\nS'il vous plaît entrez 'o' ou 'n' pour répondre.\n\n")
					password, err = generateur.Generer(reader)
					if err != nil {
						return err
					}
					fmt.Println("\nVoici votre nouveau mot de passe:", password)
					return nil
				}

				if language == "italiano" {

					fmt.Print("Benvenuto al vostro Generatore di Password.\n\nCome vorreste la vostra password?\nPer favore entra 's' o 'n' per rispondere.\n\n")

					password, err = generatore.Generare(reader)
					if err != nil {
						return err
					}
					fmt.Println("\nEcco la vostra password:", password)
					return nil
				}

				fmt.Print("Welcome to your Password Generator.\n\nHow would you like your password?\nPlease type 'y' or 'n' to give your responses.\n\n")
				password, err = generator.Generate(reader)
				if err != nil {
					return err
				}
				fmt.Println("\nBehold your new password:", password)
				return nil
			},
		},
	}
}

func main() {
	info()
	flags()
	commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
