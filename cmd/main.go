package main

import (
	"fmt"
	"github.com/MehdiEidi/goshell/commands"
	"github.com/MehdiEidi/goshell/utils"
	"log"
	"os"
)

const currentPath = "/home/mehdi/Workspace"

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	for {
		fmt.Print(wd + " :$> ")

		in := utils.GetIn()

		// && as the last field in input, represents that the new process must run concurrently.
		concurrent := in[len(in)-1] == "&&"

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
