package main

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

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
