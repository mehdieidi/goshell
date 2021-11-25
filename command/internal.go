// Package command executes various shell commands like cd.
package command

import (
	"fmt"
	"os"

	"github.com/MehdiEidi/goshell/config"
)

// CD is the implementation of the famous cd command. Returns changed working directory and error, if any.
func CD(in []string, wd string) (string, error) {
	// empty cd command or ~ indicates home
	if len(in) == 1 || in[1] == "~" {
		h, _ := os.UserHomeDir()
		err := os.Chdir(h)
		if err != nil {
			return wd, err
		}
		return h, nil
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

// Help displays a help text which explains the internal commands of the shell.
func Help() {
	fmt.Println("goshell\n-------\ncd -> change directory\nexit -> exit shell\nhelp -> display help\nabout -> display about")
}

// About just prints a nice info about the shell and authors.
func About() {
	fmt.Println(`
  xxxxxx xxxxx     xxx x  x xxx x   x
  x      x   x     x   x  x x   x   x
  x  xxx x   x     xxx xxxx xxx x   x
  x    x x   x       x x  x x   x   x
  xxxxxx xxxxx     xxx x  x xxx xxx xxx
  ─────────────────────────────────────
`, config.YELLOW, `
  Authors:`, config.CYAN, `
  * Mehdi Eidi (mehdiadq@gmail.com)
  * Reyhaneh MehdiGholizadeh (rmg62333@gmail.com)
  `, config.RESET)
}
