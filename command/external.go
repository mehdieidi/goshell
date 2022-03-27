package command

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// Cmd executes the command in[0]. in[1:]... are args.
// Concurrent says whether we must block the parent process until child process joins it, or it must continue concurrently.
func Cmd(in []string, concurrent bool) error {
	cmd := exec.Command(in[0], in[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	var err error

	if concurrent {
		err = cmd.Start()
		go cmd.Wait() // to avoid leaving zombies
	} else {
		err = cmd.Run()
	}

	return err
}

// CmdRedirect executes the command in[0] with args in[1:]... and redirects the output to the file specified after > operator.
func CmdRedirect(in []string, output bool) error {
	filename := in[len(in)-1]
	in = cleanUp(in)

	cmd := exec.Command(in[0], in[1:]...)

	if output {
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
		if err != nil {
			return fmt.Errorf("redirect openFile failed: %w", err)
		}
		defer file.Close()

		cmd.Stdin = os.Stdin
		cmd.Stdout = file

		if err := cmd.Run(); err != nil {
			return err
		}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("redirect open() failed: %w", err)
		}
		defer file.Close()

		cmd.Stdin = file
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	return nil
}

// CmdPipe executes command in[0] with args in[1:]... and pipes its output to receiver command (in[len(in)-1]).
func CmdPipe(in []string) error {
	rcv, in := extractRcv(in)

	srcCmd := exec.Command(in[0], in[1:]...)
	rcvCmd := exec.Command(rcv[0], rcv[1:]...)
	rcvCmd.Stdout = os.Stdout
	srcCmd.Stdin = os.Stdin

	r, w := io.Pipe()
	srcCmd.Stdout = w
	rcvCmd.Stdin = r

	if err := srcCmd.Start(); err != nil {
		return err
	}
	if err := rcvCmd.Start(); err != nil {
		return err
	}
	if err := srcCmd.Wait(); err != nil {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	if err := rcvCmd.Wait(); err != nil {
		return err
	}

	return nil
}

// cleanUp deletes suffix operator and operand. operators available: >, <, |
func cleanUp(in []string) []string {
	return in[:len(in)-2]
}

// extractRcv parses the input slice, extracts the receiver command and cleans the in slice.
// returns receiver command and cleaned input.
func extractRcv(in []string) ([]string, []string) {
	var r []string

	for i, v := range in {
		if v == "|" {
			r = in[i+1:]
			in = in[:i]
			break
		}
	}

	return r, in
}
