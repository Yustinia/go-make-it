package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	language := validateArgs()

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	template, ok := templates[language]
	if !ok {
		fmt.Printf("Unknown language: %v", language)
		os.Exit(1)
	}

	createTemplate(cwd, template)
}

func validateArgs() string {
	if len(os.Args) != 2 {
		fmt.Println("Usage: gmk <language>")
		os.Exit(1)
	}

	return os.Args[1]
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
