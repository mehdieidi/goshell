// Package utils contains some utility functions.
package utils

import (
	"bufio"
	"os"
	"strings"
)

// Contains returns true if s contains k
func Contains(s []string, k string) bool {
	for _, v := range s {
		if v == k {
			return true
		}
	}
	return false
}

// CleanupIn deletes suffix operator and operand.
// operators available: >, <, |
func CleanupIn(in []string) []string {
	return in[:len(in)-2]
}

// GetIn gets input, parses it, joins fields of input into a slice. in[0] is command, rest are args.
func GetIn() []string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return strings.Fields(scanner.Text())
}
