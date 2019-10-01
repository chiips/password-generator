package generator

import (
	"bufio"
	"strings"
	"testing"
	"unicode"
)

func TestGetConditions(t *testing.T) {
	upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	digits := "1234567890"
	special := "~!@#$%^&*()_+<>?"
	testAll := upper + lower + digits + special

	conditions := 4
	data := ""
	for i := 1; i <= conditions; i++ {
		data += "y\n"
	}

	reader := bufio.NewReader(strings.NewReader(data))

	password, all, conditionCount, err := getConditions(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if !passwordHasConditions(password) {
		t.Error("got", password, "want", "match all four conditions")
		t.FailNow()
	}

	if all != testAll {
		t.Error("got", all, "want", testAll)
		t.FailNow()
	}

	if conditionCount != 4 {
		t.Error("got", conditionCount, "want", conditions)
		t.FailNow()
	}

}

func passwordHasConditions(s string) bool {
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSpecial
}

func TestGetUpperTrue(t *testing.T) {

	data := "y\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsUpper, err := getUpper(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsUpper == false {
		t.Error("got", needsUpper, "want", true)
		t.FailNow()
	}

}

func TestGetUpperFalse(t *testing.T) {

	data := "n\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsUpper, err := getUpper(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsUpper == true {
		t.Error("got", needsUpper, "want", true)
		t.FailNow()
	}

}

func TestGetLowerTrue(t *testing.T) {

	data := "y\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsLower, err := getLower(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsLower == false {
		t.Error("got", needsLower, "want", true)
		t.FailNow()
	}

}

func TestGetLowerFalse(t *testing.T) {

	data := "n\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsLower, err := getLower(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsLower == true {
		t.Error("got", needsLower, "want", true)
		t.FailNow()
	}

}

func TestGetDigitsTrue(t *testing.T) {

	data := "y\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsDigits, err := getDigits(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsDigits == false {
		t.Error("got", needsDigits, "want", true)
		t.FailNow()
	}

}

func TestGetDigitsFalse(t *testing.T) {

	data := "n\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsDigits, err := getDigits(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsDigits == true {
		t.Error("got", needsDigits, "want", true)
		t.FailNow()
	}

}

func TestGetSpecialTrue(t *testing.T) {

	data := "y\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsSpecial, err := getSpecial(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsSpecial == false {
		t.Error("got", needsSpecial, "want", true)
		t.FailNow()
	}

}

func TestGetSpecialFalse(t *testing.T) {

	data := "n\n"
	reader := bufio.NewReader(strings.NewReader(data))

	needsSpecial, err := getSpecial(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if needsSpecial == true {
		t.Error("got", needsSpecial, "want", true)
		t.FailNow()
	}

}

func TestGetLength(t *testing.T) {

	data := "5\n"
	reader := bufio.NewReader(strings.NewReader(data))

	length, err := getLength(reader)

	if err != nil {
		t.Error(err.Error())
	}

	if length < 1 {
		t.Error("got", length, "want", data)
		t.FailNow()
	}

}
