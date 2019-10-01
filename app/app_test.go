package app

import (
	"os"
	"testing"
)

//TestDefaultLanguage
func TestDefaultLanguage(t *testing.T) {

	cli := NewApp()

	cli.SetInfo()
	cli.SetFlags()
	cli.SetCommands()

	expectedLang := "english"

	os.Args = []string{"password", "gen"}

	//ignore errors since focus is language flag
	_ = cli.Run(os.Args)

	//package-level language variable should equal expected
	if language != expectedLang {
		t.Error("got", language, "want", expectedLang)
	}

}

// TestAvailableLanguage
func TestAvailableLanguage(t *testing.T) {

	cli := NewApp()

	cli.SetInfo()
	cli.SetFlags()
	cli.SetCommands()

	expectedLang := "francais"

	os.Args = []string{"password", "gen", "-l=francais"}

	//ignore errors since focus is language flag
	_ = cli.Run(os.Args)

	//package-level language variable should equal expected
	if language != expectedLang {
		t.Error("got", language, "want", expectedLang)
	}

}

// TestUnavailableLanguage
func TestUnavailableLanguage(t *testing.T) {

	cli := NewApp()

	cli.SetInfo()
	cli.SetFlags()
	cli.SetCommands()

	expectedUnavailable := true

	os.Args = []string{"password", "gen", "-l=espanol"}

	//ignore errors since focus is language flag
	_ = cli.Run(os.Args)

	//package-level language variable should equal expected
	if unavailable != expectedUnavailable {
		t.Error("got", unavailable, "want", expectedUnavailable)
	}

}
