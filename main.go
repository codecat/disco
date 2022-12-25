package main

import (
	"os"
	"strings"

	"github.com/pelletier/go-toml"
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
				println("Usage: disco [--setup|-s] [type]")
				return
			} else if arg == "--setup" || arg == "-s" {
				argSetup = true
				continue
			} else {
				println("Unexpected argument:", arg)
				return
			}
		}

		if cfg.Type == "" {
			cfg.Type = arg
		} else {
			if len(cfg.Execute) > 0 {
				cfg.Execute += " "
			}
			cfg.Execute += arg
		}
	}

	if argSetup {
		setup(cfg)
	}
	start(cfg)
}
