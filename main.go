package main

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func main() {
	argSetup := false

	cfg := &DiscoConfig{}

	f, err := os.Open("disco.toml")
	if err == nil {
		defer f.Close()

		err = toml.NewDecoder(f).Decode(cfg)
		if err != nil {
			panic(err)
		}
	}

	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "-") {
			if arg == "--help" || arg == "-h" {
				println("Usage: disco [--setup|-s] [--ssh] [--zshrc] [type] [command] [args..]")
				return

			} else if arg == "--setup" || arg == "-s" {
				argSetup = true
				continue

			} else if arg == "--ssh" {
				cfg.SSH = true
				continue

			} else if arg == "--zshrc" {
				cfg.Zshrc = true

			} else {
				println("Unexpected argument:", arg)
				return
			}
		}

		if cfg.Type == "" {
			cfg.Type = arg
		} else {
			if cfg.Execute != "" {
				cfg.Execute += " "
			}
			cfg.Execute += arg
		}
	}

	if argSetup {
		setup(cfg)
		return
	}

	start(cfg)
}

func setup(cfg *DiscoConfig) {
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
