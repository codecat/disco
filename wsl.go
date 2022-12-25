package main

import (
	"regexp"
	"strings"
)

var gRegexWSLPath = regexp.MustCompile(`^([a-zA-Z]+):\/`)

func toWslPath(path string) string {
	path = strings.ReplaceAll(path, "\\", "/")
	disk := gRegexWSLPath.FindStringSubmatch(path)
	diskLetters := strings.ToLower(disk[1])
	return gRegexWSLPath.ReplaceAllString(path, "/mnt/"+diskLetters+"/")
}
