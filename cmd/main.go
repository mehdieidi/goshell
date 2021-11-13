package main

import (
	"fmt"
	"log"
	"os"
	user2 "os/user"

	"github.com/MehdiEidi/goshell/exec"
	"github.com/MehdiEidi/goshell/utils"
)

func main() {
	// current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	// current user
	user, err := user2.Current()
	if err != nil {
		fmt.Println(err)
	}

	// machines name
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(host)
	}

	for {
		fmt.Print(red, user.Username+"@"+host+" ", blue, wd, yellow, " >>> ", reset)

		// get input command and args in a slice
		in := utils.GetIn()

		// && at the end of the input represents that the new process must run concurrently with parent process.
		var concurrent bool
		if in[len(in)-1] == "&&" {
			concurrent = true
			in = in[:len(in)-1] // delete &&
		}

		switch in[0] {
		case "cd":
			wd = exec.CD(in, wd)

		case "exit":
			exec.Exit()

		default:
			switch {
			case utils.Contains(in, ">"):
				ExecRedirect(in)

			case utils.Contains(in, "|"):
				ExecPipe(in)

			default:
				Exec(in, concurrent)
			}
		}
	}
}
