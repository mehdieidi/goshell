package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Exec executes the command in[0]. in[1]... are args.
// concurrent says whether we must block the parent process until child process joins it, or it must continue concurrently.
func Exec(in []string, concurrent bool) {
	cmd := exec.Command(in[0], in[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	var err error

	if concurrent {
		err = cmd.Start()
	} else {
		err = cmd.Run()
	}

	if err != nil {
		fmt.Println(err)
	}
}
