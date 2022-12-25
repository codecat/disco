package main

type DiscoConfig struct {
	Type    string `toml:"type"`
	Execute string `toml:"execute,omitempty"`
	Pull    bool   `toml:"pull,omitempty"`

	Ports   []string `toml:"ports,omitempty"`
	Network string   `toml:"network,omitempty"`
	Options string   `toml:"options,omitempty"`

	SSH   bool `toml:"ssh,omitempty"`
	Zshrc bool `toml:"zshrc,omitempty"`
}
