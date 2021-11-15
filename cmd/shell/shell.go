// Package shell starts the shell and gets commands.
package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MehdiEidi/goshell/config"
	"github.com/MehdiEidi/goshell/run"
)

// Start gets the config file, runs the shell and gets commands.
func Start(c config.Config) {
	var latestCmd []string

	for {
		fmt.Print(c.UserColor, c.User.Username+"@"+c.Hostname+" ", c.PathColor, c.WD, c.PromptColor, " >>> ", c.ResetColor)

		input := getIn()
		if len(input) == 0 || input[0] == "" {
			continue
		}

		// handle history feature; !! means give me the latest cmd.
		if input[0] != "!!" {
			latestCmd = input
		} else {
			if len(latestCmd) == 0 {
				fmt.Println("History is empty...")
				continue
			} else {
				fmt.Println("Command ", latestCmd, " ran from history...")
				fmt.Println("-------------------------------------------")
				input = latestCmd
			}
		}

		input, concurrent := isConcurrent(input)

		switch input[0] {
		case "cd":
			w, err := run.CD(input, c.WD)
			if err != nil {
				fmt.Println(err)
			}
			c.WD = w

		case "exit":
			run.Exit()

		default:
			switch {
			case contains(input, ">"):
				err := run.CmdRedirect(input, true)
				if err != nil {
					fmt.Println(err)
				}

			case contains(input, "<"):
				err := run.CmdRedirect(input, false)
				if err != nil {
					fmt.Println(err)
				}

			case contains(input, "|"):
				err := run.CmdPipe(input)
				if err != nil {
					fmt.Println(err)
				}

			default:
				err := run.Cmd(input, concurrent)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// getIn gets input, parses it, joins fields of input into a slice.
func getIn() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.Split(scanner.Text(), " ")
}

// isConcurrent returns true if the command in[0] is supposed to run concurrently with parent.
// && at the end of the command means that the command must run concurrently.
// It also cleans the && from the in slice and returns the clean slice.
func isConcurrent(in []string) ([]string, bool) {
	if in[len(in)-1] == "&&" {
		in = in[:len(in)-1] // delete &&
		return in, true
	}
	return in, false
}

// contains returns true if s contains k
func contains(s []string, k string) bool {
	for _, v := range s {
		if v == k {
			return true
		}
	}
	return false
}
