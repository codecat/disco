package main

import (
	"os"
	"regexp"
	"runtime"
	"strings"

	"github.com/pelletier/go-toml"
)

var gRegexWSLPath = regexp.MustCompile(`^([a-zA-Z]+):\/`)

type DiscoConfig struct {
	Type  string
	Ports []string
}

func main() {
	for _, arg := range os.Args[1:] {
		if arg == "--help" || arg == "-h" {
			println("Usage: disco")
			return
		}
	}

	f, err := os.Open("disco.toml")
	if err != nil {
		println(err.Error())
		return
	}

	cfg := DiscoConfig{}
	err = toml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		println(err.Error())
		return
	}

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
	flags += " -v \"" + workdir + "/disco.toml:/src/disco.toml:ro\""

	switch cfg.Type {
	case "js":
	case "javascript":
		image = "codecatt/disco:js"

	case "vite":
		image = "codecatt/disco:js"
		flags += " -p 127.0.0.1:5173:5173"

	case "py":
	case "python":
		image = "codecatt/disco:py"
	}

	for _, port := range cfg.Ports {
		flags += " -p " + port
	}

	cmd += "docker run " + flags + " " + image
	println(cmd)
}

func toWslPath(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	disk := gRegexWSLPath.FindStringSubmatch(path)
	diskLetters := strings.ToLower(disk[1])
	return gRegexWSLPath.ReplaceAllString(path, "/mnt/"+diskLetters+"/")
}
