package main

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func main() {
	argInit := false

	cfg := &DiscoConfig{}

	f, err := os.Open("disco.toml")
	if err == nil {
		defer f.Close()

		err = toml.NewDecoder(f).Decode(cfg)
		if err != nil {
			panic(err)
		}
	}

	readingDiscoArgs := true
	for _, arg := range os.Args[1:] {
		if readingDiscoArgs && strings.HasPrefix(arg, "-") {
			if arg == "--help" || arg == "-h" {
				println("Usage: disco [--init|-i] [--build|-b] [--ssh] [--fish] [type] [command] [args..]")
				return

			} else if arg == "--init" || arg == "-i" {
				argInit = true
				continue

			} else if arg == "--build" || arg == "-b" {
				cfg.Build = true
				continue

			} else if arg == "--ssh" {
				cfg.SSH = true
				continue

			} else if arg == "--fish" {
				cfg.Fish = true
				continue

			} else {
				println("Unexpected argument:", arg)
				return
			}
		}

		readingDiscoArgs = false

		if cfg.Type == "" {
			cfg.Type = arg
			continue
		}

		if cfg.Execute != "" {
			cfg.Execute += " "
		}
		cfg.Execute += arg
	}

	if argInit {
		initConfig(cfg)
		return
	}

	start(cfg)
}

func initConfig(cfg *DiscoConfig) {
	if cfg.Type == "" {
		cfg.Type = "base"
	}

	_, err := os.Stat("disco.toml")
	if err == nil {
		println("disco.toml already exists")
		return
	}

	f, err := os.Create("disco.toml")
	if err != nil {
		panic(err)
	}
	toml.NewEncoder(f).Encode(cfg)
	f.Close()

	println("Created disco.toml")
}
