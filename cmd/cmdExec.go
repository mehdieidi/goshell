package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/MehdiEidi/goshell/utils"
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

// ExecRedirect executes the command in[0] with args in[1]...
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

	file, err := os.OpenFile(currentPath+"/"+filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := io.Copy(io.MultiWriter(file, os.Stdout), cmdStdout); err != nil {
		fmt.Println(err)
	}

	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}
}

//func ExecPipe(in []string) {
//	receiver := in[len(in)-1]
//	in = utils.CleanupIn(in)
//
//	srcCmd := exec.Command(in[0], in[1:]...)
//	stdout, err := srcCmd.StdoutPipe()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	rcvCmd := exec.Command(receiver)
//	stdin, err := rcvCmd.StdinPipe()
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	if err = srcCmd.Run(); err != nil {
//		fmt.Println(err)
//	}
//	if _, err = io.Copy(stdin, stdout); err != nil {
//		fmt.Println(err)
//	}
//	if err = rcvCmd.Run(); err != nil {
//		fmt.Println(err)
//	}
//}

func ExecPipe(in []string) {
	receiver := in[len(in)-1]
	in = utils.CleanupIn(in)

	srcCmd := exec.Command(in[0], in[1:]...)
	//stdout, err := srcCmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println(err)
	//}

	rcvCmd := exec.Command(receiver)
	stdin, err := rcvCmd.StdinPipe()
	if err != nil {
		fmt.Println(err)
	}

	srcCmd.Start()
	str, _ := srcCmd.Output()
	fmt.Fprintln(stdin, string(str))

	rcvCmd.Start()
	str, _ = rcvCmd.Output()
	fmt.Fprintln(os.Stdout, string(str))

	srcCmd.Wait()
	//rcvCmd.Wait()

}