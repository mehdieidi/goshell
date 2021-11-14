// Package config initializes some data for shell.
package config

import (
	"os"
	userPkg "os/user"
)

// Config contains some details about the host which will be used
// in shell.
type Config struct {
	WD          string
	Host        string
	User        *userPkg.User
	UserColor   string
	PathColor   string
	PromptColor string
	ResetColor  string
}

// New returns a new Config value initialized with correct data. It also returns errors if any.
func New() (Config, error) {
	// current working directory
	wd, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}

	// current User
	user, err := userPkg.Current()
	if err != nil {
		return Config{}, err
	}

	// Host
	host, err := os.Hostname()
	if err != nil {
		return Config{}, err
	}

	c := Config{
		WD:          wd,
		Host:        host,
		User:        user,
		UserColor:   red,
		PathColor:   blue,
		PromptColor: yellow,
		ResetColor:  reset,
	}

	return c, nil
}
