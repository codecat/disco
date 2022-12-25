package main

import (
	"fmt"
	"os"
)

func setup(cfg *DiscoConfig) {
	if cfg.Type == "" {
		panic("Missing type!")
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
	defer f.Close()

	fmt.Fprintf(f, "type = \"%s\"\n", cfg.Type)
	println("Created disco.toml")
}
