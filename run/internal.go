// Package run runs various shell commands like cd
package run

import (
	"os"
)

// CD is the implementation of the famous cd command. Returns changed working directory.
func CD(in []string, wd string) (string, error) {
	// empty cd command or ~ indicates home
	if len(in) == 1 || in[1] == "~" {
		h, _ := os.UserHomeDir()
		err := os.Chdir(h)
		if err != nil {
			return wd, err
		}
		return h, err
	}

	// Change working directory
	err := os.Chdir(in[1])
	if err != nil {
		return wd, err
	} else {
		wd, _ = os.Getwd()
	}

	return wd, nil
}

// Exit just exits the program with exit code 0.
func Exit() {
	os.Exit(0)
}
