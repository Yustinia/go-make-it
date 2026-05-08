package main

import "fmt"

func printHelp() {
	fmt.Println("Usage: gmk <language>")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println("  -h, --help    Print usage and available flags")
	fmt.Println("  -l, --list    List supported languages")
}

func printList() {
	fmt.Println("Supported languages:")
	for key := range templates {
		fmt.Println(" ", key)
	}
}
