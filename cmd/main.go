package main

import (
	"github.com/MehdiEidi/goshell/cmd/shell"
	"github.com/MehdiEidi/goshell/config"
)

func main() {
	c, err := config.New()
	if err != nil {
		panic(err)
	}

	shell.Start(c)
}
