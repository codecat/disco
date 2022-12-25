package main

type DiscoConfig struct {
	Type    string `toml:"type"`
	Execute string `toml:"execute,omitempty"`

	Ports []string `toml:"ports,omitempty"`

	SSH bool `toml:"ssh,omitempty"`
}
