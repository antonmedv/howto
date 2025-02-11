//go:build !windows

package main

import (
	"os"
	"syscall"
	"unsafe"

	"golang.org/x/term"
)

func insertInput(cmd string) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for _, c := range cmd {
		char := byte(c)
		syscall.Syscall(syscall.SYS_IOCTL, uintptr(0), syscall.TIOCSTI, uintptr(unsafe.Pointer(&char)))
	}
}
