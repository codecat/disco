package main

import (
	"os"
	"os/exec"
	"runtime"
)

func start(cfg *DiscoConfig) {
	cmd := ""
	workdir, _ := os.Getwd()
	homedir, _ := os.UserHomeDir()

	if runtime.GOOS == "windows" {
		cmd = "wsl --exec "
		workdir = toWslPath(workdir)
		homedir = toWslPath(homedir)
	}

	image := "codecatt/disco:base"
	flags := "--rm -it"
	flags += " -v \"" + homedir + "/.ssh:/home/developer/.ssh:ro\""
	flags += " -v \"" + workdir + ":/src\""

	switch cfg.Type {
	case "base":
		// Don't have to change anything

	case "js", "javascript":
		image = "codecatt/disco:js"

	case "vite":
		image = "codecatt/disco:js"
		flags += " -p 127.0.0.1:5173:5173"

	case "py", "python":
		image = "codecatt/disco:py"
	}

	for _, port := range cfg.Ports {
		flags += " -p " + port
	}

	cmd += "docker run " + flags + " " + image

	if cfg.Execute != "" {
		cmd += " " + cfg.Execute
	}

	c := exec.Command("sh", "-c", cmd)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Run()
}
