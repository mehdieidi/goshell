// Package shell starts the shell.
package shell

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MehdiEidi/goshell/config"
	"github.com/MehdiEidi/goshell/run"
)

// Start gets the config file and runs the shell.
func Start(c config.Config) {
	for {
		fmt.Print(c.UserColor, c.User.Username+"@"+c.Host+" ", c.PathColor, c.WD, c.PromptColor, " >>> ", c.ResetColor)

		in := getIn()
		if len(in) == 0 {
			continue
		}

		concurrent := isConcurrent(in)

		switch in[0] {
		case "cd":
			c.WD = run.CD(in, c.WD)

		case "exit":
			run.Exit()

		default:
			switch {
			case contains(in, ">"):
				err := run.CmdRedirect(in, true)
				if err != nil {
					fmt.Println(err)
				}

			case contains(in, "<"):
				err := run.CmdRedirect(in, false)
				if err != nil {
					fmt.Println(err)
				}

			case contains(in, "|"):
				err := run.CmdPipe(in)
				if err != nil {
					fmt.Println(err)
				}

			default:
				err := run.Cmd(in, concurrent)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// getIn gets input, parses it, joins fields of input into a slice. in[0] is command, rest are args.
func getIn() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.Fields(scanner.Text())
}

// isConcurrent returns true if the command in[0] supposed to run
// concurrently with parent. && at the end of the command means that the command
// must run concurrently. It also cleans the && from the in slice.
func isConcurrent(in []string) bool {
	if in[len(in)-1] == "&&" {
		in = in[:len(in)-1] // delete &&
		return true
	}
	return false
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
