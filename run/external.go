package run

import (
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
		cmdStdout, err := cmd.StdoutPipe()
		if err != nil {
			return err
		}

		file, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
		if err != nil {
			return err
		}
		defer file.Close()

		if err := cmd.Start(); err != nil {
			return err
		}

		if _, err := io.Copy(file, cmdStdout); err != nil {
			return err
		}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			return err
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
	receiver := in[len(in)-1]
	in = cleanUp(in)

	srcCmd := exec.Command(in[0], in[1:]...)
	rcvCmd := exec.Command(receiver)
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
