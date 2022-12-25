package main

import (
	"os"
	"os/exec"
	"runtime"
)

func getDockerCommand() string {
	if runtime.GOOS == "windows" {
		return "wsl --exec docker"
	}
	return "docker"
}

func newDockerCommand(args string) *exec.Cmd {
	exe := "sh"
	exeFlag := "-c"

	if runtime.GOOS == "windows" {
		exe = "cmd"
		exeFlag = "/c"
	}

	return exec.Command(exe, exeFlag, getDockerCommand()+" "+args)
}

func runDocker(args string) error {
	cmd := newDockerCommand(args)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
