// Package run executes various shell commands like cd.
package run

import (
	"fmt"
	"os"
)

// CD is the implementation of the famous cd command. Returns changed working directory and error, if any.
func CD(in []string, wd string) (string, error) {
	// empty cd command or ~ indicates home
	if len(in) == 1 || in[1] == "~" {
		h, _ := os.UserHomeDir()
		err := os.Chdir(h)
		if err != nil {
			return wd, fmt.Errorf("chdir() failed: %w", err)
		}
		return h, nil
	}

	// Change working directory
	err := os.Chdir(in[1])
	if err != nil {
		return wd, fmt.Errorf("chdir() failed: %w", err)
	} else {
		wd, _ = os.Getwd()
	}

	return wd, nil
}

// Exit just exits the program with exit code 0.
func Exit() {
	os.Exit(0)
}

// Help displays a help text which explains the internal commands of the shell.
func Help() {
	fmt.Println("cd -> change directory\nexit -> exit shell\nhelp -> display help")
}
