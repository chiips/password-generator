package generateur

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestGenerer(t *testing.T) {

	//set responses
	conditions := 4
	data := ""
	for i := 1; i <= conditions; i++ {
		data += "y\n"
	}

	length := 5
	data += fmt.Sprintf("%d\n", length)

	reader := bufio.NewReader(strings.NewReader(data))

	password, err := Generer(reader)

	if err != nil {
		t.Error(err.Error())
	}

	//passwordHasConditions declared in conditions_test.go
	if !passwordHasConditions(password) {
		t.Error("got", password, "want", "match all four conditions")
		t.FailNow()
	}

	if len(password) != length {
		t.Error("got", length, "want", data)
		t.FailNow()
	}

}
