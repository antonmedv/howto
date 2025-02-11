package main

import (
	"runtime"
)

func userOs() string {
	goos := runtime.GOOS
	switch goos {
	case "darwin":
		return "mac"
	}
	return goos
}
