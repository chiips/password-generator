package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/chiips/password/generator"
)

func main() {

	fmt.Printf("Welcome to your Password Generator.\n\nHow would you like your password?\nPlease type 'y' or 'n' to give your responses.\n\n")

	reader := bufio.NewReader(os.Stdin)
	password, err := generator.Generate(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println("Behold your new password:", password)

}