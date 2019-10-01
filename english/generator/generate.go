package generator

import (
	"bufio"
	"crypto/rand"
	"math/big"
)

//Generate creates the password based on user input
func Generate(reader *bufio.Reader) (string, error) {

	//get iniital password with guaranteed conditions, our total character selection, and the count of conditions selected
	password, all, conditionCount, err := getConditions(reader)
	if err != nil {
		return "", err
	}

	//get request length
	length, err := getLength(reader)
	if err != nil {
		return "", err
	}

	//subtract condition count from length because that many characters already added
	//then for the length remaining randomly add to the password from our total character selection.
	for i := 0; i < length-conditionCount; i++ {

		max := big.NewInt(int64(len(all)))

		index, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		char := string(all[index.Int64()])

		password = insert(password, char)

	}

	//ensure password length matches requested length.
	//only for cases where all 4 conditions selected but length < 4 requested
	if len(password) > length {
		password = password[0:length]
	}

	return password, nil

}
