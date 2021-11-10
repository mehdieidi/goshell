package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	for {
		fmt.Print(wd + " :$> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		in := strings.Fields(strings.TrimSuffix(scanner.Text(), "\n"))

		concurrent := in[len(in)-1] == "&&"

		switch in[0] {
		case "cd":
			wd = cd(in, wd)

		case "exit":
			os.Exit(0)

		default:
			cmd := exec.Command(in[0], in[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if concurrent {
				err = cmd.Start()
				if err != nil {
					fmt.Println(err)
				}
			} else {
				err = cmd.Run()
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

// cd is the handler of the famous cd command.
// cd supports: absolute path, relative path
func cd(in []string, wd string) string {
	var path string

	// pure cd command
	if len(in) == 1 {
		path = "/home/mehdi"
	} else if len(in) > 1 {
		if in[1] == ".." {
			pf := strings.Split(wd, "/")
			pf = pf[:len(pf)-1]
			path = strings.Join(pf, "/")
		} else if in[1] == "~" {
			path = "/home/mehdi"
		} else {
			// relative path
			if in[1][0] != '/' {
				path = wd + "/" + in[1]
			} else {
				path = in[1]
			}
		}
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Println(err)
	} else {
		wd = path
	}

	return wd
}
