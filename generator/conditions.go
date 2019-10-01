package generator

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
	digits := "1234567890"
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
	needDigits, err := getDigits(reader)
	if err != nil {
		return "", "", 0, err
	}
	if needDigits {
		max := big.NewInt(int64(len(digits)))

		i, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", "", 0, err
		}

		char := string(digits[i.Int64()])

		password = insert(password, char)

		all += digits
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
		fmt.Println("You must select at least one condition")
		return getConditions(reader)
	}

	return password, all, conditionCount, nil
}

func getUpper(reader *bufio.Reader) (bool, error) {
	fmt.Printf("With upper case letters? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)

	if answer == "y" || answer == "Y" {
		return true, nil
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Please enter 'y' or 'n' for your response.")
	return getUpper(reader)

}

func getLower(reader *bufio.Reader) (bool, error) {
	fmt.Printf("With lower case letters? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)
	if answer == "y" || answer == "Y" {
		return true, nil
	}
	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Please enter 'y' or 'n' for your response.")
	return getLower(reader)

}

func getDigits(reader *bufio.Reader) (bool, error) {
	fmt.Printf("With digits? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)

	if answer == "y" || answer == "Y" {
		return true, nil
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Please enter 'y' or 'n' for your response.")
	return getDigits(reader)

}

func getSpecial(reader *bufio.Reader) (bool, error) {
	fmt.Printf("With special characters? ")
	answer, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	answer = strings.TrimSpace(answer)

	if answer == "y" || answer == "Y" {
		return true, nil
	}

	if answer == "n" || answer == "N" {
		return false, nil
	}

	fmt.Println("Please enter 'y' or 'n' for your response.")
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

	fmt.Printf("Please enter your password length: ")

	lengthStr, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	lengthTrim := strings.TrimSpace(lengthStr)

	length, err := strconv.Atoi(lengthTrim)
	if err != nil {
		fmt.Println("Please enter a number for the length.")
		return getLength(reader)

	}

	if length >= 1 {
		return length, nil
	}

	fmt.Println("Password must be at least 1 character long.")
	return getLength(reader)

}
