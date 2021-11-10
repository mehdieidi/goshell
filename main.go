package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("$> ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()

		in := strings.Fields(strings.TrimSuffix(scanner.Text(), "\n"))

		cmd := exec.Command(in[0], in[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}
