package main

import (
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
)

var gRegexWSLPath = regexp.MustCompile(`^([a-zA-Z]+):\/`)

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

func toWslPath(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	disk := gRegexWSLPath.FindStringSubmatch(path)
	diskLetters := strings.ToLower(disk[1])
	return gRegexWSLPath.ReplaceAllString(path, "/mnt/"+diskLetters+"/")
}
