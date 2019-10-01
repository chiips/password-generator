package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/chiips/password/english/generator"
	"github.com/chiips/password/francais/generateur"
	"github.com/chiips/password/italiano/generatore"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	genCommand := flag.NewFlagSet("gen", flag.ExitOnError)
	lang := genCommand.String("l", "english", "Select your language from english, francais, italiano.")

	help := "Password Generator generates cryptographically secure passwords. Choose your"

	if len(os.Args) < 2 {
		fmt.Println(help)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "gen":
		genCommand.Parse(os.Args[2:])
	default:
		fmt.Println(help)
		os.Exit(1)
	}

	options := map[string]bool{"english": true, "francais": true, "italiano": true}
	if _, valid := options[*lang]; !valid {
		genCommand.PrintDefaults()
		os.Exit(1)

	}

	var password string
	var err error

	if fmt.Sprintf("%s", *lang) == "english" {
		fmt.Printf("Welcome to your Password Generator.\n\nHow would you like your password?\nPlease type 'y' or 'n' to give your responses.\n\n")
		password, err = generator.Generate(reader)
		if err != nil {
			fmt.Println("There was error processing your password request:", err)
			os.Exit(1)
		}
	}
	if fmt.Sprintf("%s", *lang) == "francais" {
		fmt.Printf("Bienvenue to your Password Generator.\n\nHow would you like your password?\nPlease type 'y' or 'n' to give your responses.\n\n")
		password, err = generateur.Generer(reader)
		if err != nil {
			fmt.Println("There was error processing your password request:", err)
			os.Exit(1)
		}
	}

	if fmt.Sprintf("%s", *lang) == "italiano" {
		fmt.Printf("Benvenuto to your Password Generator.\n\nHow would you like your password?\nPlease type 'y' or 'n' to give your responses.\n\n")
		password, err = generatore.Generare(reader)
		if err != nil {
			fmt.Println("There was error processing your password request:", err)
			os.Exit(1)
		}
	}

	fmt.Println("\nBehold your new password:", password)

}
