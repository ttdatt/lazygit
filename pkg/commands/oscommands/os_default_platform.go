//go:build !windows
// +build !windows

package oscommands

import (
	"os"
	"runtime"
)

func GetPlatform() *Platform {
	shell := getUserShell()

	interactiveShell := shell
	interactiveShellArg := "-i"
	interactiveShellExit := "; exit"

	return &Platform{
		OS:                   runtime.GOOS,
		Shell:                "bash",
		InteractiveShell:     interactiveShell,
		ShellArg:             "-c",
		InteractiveShellArg:  interactiveShellArg,
		InteractiveShellExit: interactiveShellExit,
		OpenCommand:          "open {{filename}}",
		OpenLinkCommand:      "open {{link}}",
	}
}

func getUserShell() string {
	if shell := os.Getenv("SHELL"); shell != "" {
		return shell
	}

	return "bash"
}
