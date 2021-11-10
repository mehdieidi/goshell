package commands

import (
	"fmt"
	"os"
	"strings"
)

// CD is the implementation of famous cd command.
// returns changed working directory.
func CD(in []string, wd string) string {
	var newPath string

	switch {
	case len(in) == 1:
		newPath = "/home/mehdi"

	case len(in) > 1:
		switch {
		case in[1] == "..":
			pd := strings.Split(wd, "/")
			pd = pd[:len(pd)-1]
			newPath = strings.Join(pd, "/")

		case in[1] == "~":
			newPath = "/home/mehdi"

		default:
			// relative or absolute path
			if in[1][0] != '/' {
				newPath = wd + "/" + in[1]
			} else {
				newPath = in[1]
			}
		}
	}

	// Changing working directory
	err := os.Chdir(newPath)
	if err != nil {
		fmt.Println(err)
	} else {
		wd = newPath
	}

	return wd
}

// Exit just exits the program with exit code 0.
func Exit() {
	os.Exit(0)
}
