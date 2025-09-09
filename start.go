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

	flags := "--rm -it"
	flags += " -v \"" + workdir + ":/src\""

	if cfg.SSH {
		flags += " -v \"" + homedir + "/.ssh:/home/developer/.ssh:ro\""
	}

	if cfg.Fish {
		flags += " -v \"" + homedir + "/.config/fish:/home/developer/.config/fish:ro\""
	}

	image := "base"

	switch cfg.Type {
	case "base", "":
		image = "base"

	case "js", "javascript":
		image = "js"

	case "vite":
		image = "js"
		flags += " -p 127.0.0.1:5173:5173"

	case "py", "python":
		image = "py"

	case "php":
		image = "php"

	case "php-framework":
		image = "php-framework"

	case "php-rr":
		image = "php-framework"
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

	if image != "base" && (!imageExists(image) || cfg.Build) {
		if err := buildImage("base"); err != nil {
			return
		}
	}
	if !imageExists(image) || cfg.Build {
		if err := buildImage(image); err != nil {
			return
		}
	}

	args := "run " + flags + " " + imageName(image)

	if cfg.Execute != "" {
		args += " " + cfg.Execute
	}

	println("Run: " + args)
	runDocker(args)
}
