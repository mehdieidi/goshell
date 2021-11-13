// Package exec runs various shell commands like cd
package exec

import (
	"fmt"
	"os"
)

// CD is the implementation of the thefamous cd command.
// returns changed working directory.
func CD(in []string, wd string) string {
	// empty cd command
	if len(in) == 1 {
		in = append(in, "/home/mehdi")
	}

	// Change working directory
	err := os.Chdir(in[1])
	if err != nil {
		fmt.Println(err)
	} else {
		wd, _ = os.Getwd()
	}

	return wd
}

// Exit just exits the program with exit code 0.
func Exit() {
	os.Exit(0)
}
