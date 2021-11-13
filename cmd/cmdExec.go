package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/MehdiEidi/goshell/utils"
)

// Exec executes the command in[0]. in[1:]... are args.
// Concurrent says whether we must block the parent process until child process joins it, or it must continue concurrently.
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

// ExecRedirect executes the command in[0] with args in[1:]...
// and redirects the output to the file specified after > operator.
func ExecRedirect(in []string) {
	filename := in[len(in)-1]
	in = utils.CleanupIn(in)

	cmd := exec.Command(in[0], in[1:]...)
	cmdStdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println(err)
	}

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := io.Copy(file, cmdStdout); err != nil {
		fmt.Println(err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

// ExecPipe executes command in[0] with args in[1:]...
// and pipes its output to receiver command (in[len(in)-1]).
func ExecPipe(in []string) {
	receiver := in[len(in)-1]
	in = utils.CleanupIn(in)

	srcCmd := exec.Command(in[0], in[1:]...)
	rcvCmd := exec.Command(receiver)
	rcvCmd.Stdout = os.Stdout

	r, w := io.Pipe()
	srcCmd.Stdout = w
	rcvCmd.Stdin = r

	if err := srcCmd.Start(); err != nil {
		fmt.Println(err)
	}
	if err := rcvCmd.Start(); err != nil {
		fmt.Println(err)
	}
	if err := srcCmd.Wait(); err != nil {
		fmt.Println(err)
	}
	if err := w.Close(); err != nil {
		fmt.Println(err)
	}
	if err := rcvCmd.Wait(); err != nil {
		fmt.Println(err)
	}
}
