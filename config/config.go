// Package config initializes some data for shell. e.g. username, colors, hostname...
package config

import (
	"fmt"
	"os"
	userPkg "os/user"
)

// Config contains some info that will be used in shell.
type Config struct {
	User        *userPkg.User
	WD          string
	Hostname    string
	UserColor   string
	PathColor   string
	PromptColor string
	ResetColor  string
}

// New returns a new Config value initialized with correct data. It also returns errors, if any.
func New() (Config, error) {
	// current working directory
	wd, err := os.Getwd()
	if err != nil {
		return Config{}, fmt.Errorf("getwd() failed: %w", err)
	}

	// current User
	user, err := userPkg.Current()
	if err != nil {
		return Config{}, fmt.Errorf("user.Current() failed: %w", err)
	}

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		return Config{}, fmt.Errorf("hostname() failed: %w", err)
	}

	c := Config{
		WD:          wd,
		Hostname:    hostname,
		User:        user,
		UserColor:   RED,
		PathColor:   BLUE,
		PromptColor: YELLOW,
		ResetColor:  RESET,
	}

	return c, nil
}
