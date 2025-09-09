package main

import (
	"os"
	"runtime"
)

func start(cfg *DiscoConfig) {
	workdir, _ := os.Getwd()
	homedir, _ := os.UserHomeDir()

	if runtime.GOOS == "windows" {
		workdir = toWslPath(workdir)
		homedir = toWslPath(homedir)
	}

	image := "codecatt/disco:base"
	flags := "--rm -it"
	flags += " -v \"" + workdir + ":/src\""

	if cfg.Type != "base" && !imageExists(cfg.Type) {
		buildImage("base")
	}
	if !imageExists(cfg.Type) || cfg.Build {
		buildImage(cfg.Type)
	}

	if cfg.SSH {
		flags += " -v \"" + homedir + "/.ssh:/home/developer/.ssh:ro\""
	}

	if cfg.Fish {
		flags += " -v \"" + homedir + "/.config/fish:/home/developer/.config/fish:ro\""
	}

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

	case "php":
		image = "codecatt/disco:php"

	case "php-framework":
		image = "codecatt/disco:php-framework"

	case "php-rr":
		image = "codecatt/disco:php-framework"
		flags += " -p 127.0.0.1:8080:8080"
	}

	for _, port := range cfg.Ports {
		flags += " -p " + port
	}

	if cfg.Network != "" {
		flags += " --net " + cfg.Network
	}

	if cfg.Options != "" {
		flags += " " + cfg.Options
	}

	args := "run " + flags + " " + image

	if cfg.Execute != "" {
		args += " " + cfg.Execute
	}

	runDocker(args)
}
