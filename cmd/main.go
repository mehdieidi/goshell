package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MehdiEidi/goshell/commands"
	"github.com/MehdiEidi/goshell/utils"
)

func main() {
	// current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	for {
		fmt.Print(wd + " :$> ")

		// get input command and args in a slice
		in := utils.GetIn()

		// && at the end of input represents that the new process must run concurrently with parent process.
		var concurrent bool
		if in[len(in)-1] == "&&" {
			concurrent = true
			in = in[:len(in)-1] // delete &&
		}

		switch in[0] {
		case "cd":
			wd = commands.CD(in, wd)

		case "exit":
			commands.Exit()

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
