package generateur

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func getConditions(reader *bufio.Reader) (string, string, int, error) {

	password := ""

	//CONDITIONS
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	numbers := "1234567890"
	special := "~!@#$%^&*()_+<>?"
	all := ""
	conditionCount := 0

	//UPPER
	needUpper, err := getUpper(reader)
	if err != nil {
		return "", "", 0, err
	}
	if needUpper {
		//add a random upper character to the password
		max := big.NewInt(int64(len(upper)))

		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", "", 0, err
		}

		//ensure at least one upper is in the password string
		char := string(upper[i.Int64()])

		password = insert(password, char)

		//add the upper characters to our all conditions string
		all += upper
		//increase condition count for later index
		conditionCount++
	}

	//LOWER
	needLower, err := getLower(reader)
	if err != nil {
		return "", "", 0, err
	}
	if needLower {
		max := big.NewInt(int64(len(lower)))

		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", "", 0, err
		}

		char := string(lower[i.Int64()])

		password = insert(password, char)

		all += lower
		conditionCount++
	}

	//DIGITS
	needNumbers, err := getNumbers(reader)
	if err != nil {
		return "", "", 0, err
	}
	if needNumbers {
		max := big.NewInt(int64(len(numbers)))

		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", "", 0, err
		}

		char := string(numbers[i.Int64()])

		password = insert(password, char)

		all += numbers
		conditionCount++
	}

	//SPECIAL
	needSpecial, err := getSpecial(reader)
	if err != nil {
		return "", "", 0, err
	}
	if needSpecial {
		max := big.NewInt(int64(len(special)))

		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", "", 0, err
		}

		char := string(special[i.Int64()])

		password = insert(password, char)

		all += special
		conditionCount++
	}

	if conditionCount == 0 {
		fmt.Println("Il faut sélectionner au moins une condition.")
		return getConditions(reader)
	}

	return password, all, conditionCount, nil
}

func getUpper(reader *bufio.Reader) (bool, error) {
	fmt.Printf("Avec des lettres capitales? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)

	if answer == "o" || answer == "O" {
		return true, nil
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Veuillez taper 'o' ou 'n' pour répondre.")
	return getUpper(reader)

}

func getLower(reader *bufio.Reader) (bool, error) {
	fmt.Printf("Avec des lettres minuscules? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)
	if answer == "o" || answer == "O" {
		return true, nil
	}
	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Veuillez taper 'o' ou 'n' pour répondre.")
	return getLower(reader)

}

func getNumbers(reader *bufio.Reader) (bool, error) {
	fmt.Printf("Avec des chiffres? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)

	if answer == "o" || answer == "O" {
		return true, nil
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Veuillez taper 'o' or 'n' pour répondre.")
	return getNumbers(reader)

}

func getSpecial(reader *bufio.Reader) (bool, error) {
	fmt.Printf("Avec des caractères spéciaux? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)

	if answer == "o" || answer == "O" {
		return true, nil
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Veuillez taper 'o' or 'n' pour répondre.")
	return getSpecial(reader)

}

func insert(password, char string) string {

	if password == "" {
		return char
	}

	max := big.NewInt(int64(len(password)))

	index, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}

	//randomly add to password from total selection
	password = password[0:index.Int64()] + char + password[index.Int64():len(password)]

	return password

}

func getLength(reader *bufio.Reader) (int, error) {

	fmt.Printf("Veuillez taper la longueur de votre mot de passe: ")

	lengthStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	lengthTrim := strings.TrimSpace(lengthStr)

	length, err := strconv.Atoi(lengthTrim)
	if err != nil {
		fmt.Println("Veuillez taper un chiffre pour la longueur.")
		return getLength(reader)

	}

	if length >= 1 {
		return length, nil
	}

	fmt.Println("Il faut que votre mot de passe comporte au moins un caractère.")
	return getLength(reader)

}
