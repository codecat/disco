package main

import "embed"

//go:embed images/*
var gImagesFS embed.FS

func imageName(imageType string) string {
	return "codecatt/disco:" + imageType
}

func imageExists(imageType string) bool {
	args := "images -q " + imageName(imageType)
	output, err := runDockerOutput(args)
	return err == nil && output != ""
}

func buildImage(imageType string) error {
	dockerfile, err := gImagesFS.ReadFile("images/" + imageType + ".dockerfile")
	if err != nil {
		return err
	}
	return runDockerInput("build -t "+imageName(imageType)+" -", dockerfile)
}
