package main

import (
	"os"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	createTemplate(cwd, templates["python"])
}

func createTemplate(cwd string, template []string) {
	for _, path := range template {
		fullPath := filepath.Join(cwd, path)

		err := os.MkdirAll(fullPath, 0755)
		if err != nil {
			panic(err)
		}
	}
}
