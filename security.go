package main

import (
	"strings"
)

func sanitizeCommand(cmd string) string {
	cmd = strings.TrimSpace(cmd)
	cmd = strings.Trim(cmd, "`")
	return strings.ReplaceAll(cmd, "\n", " ")
}
