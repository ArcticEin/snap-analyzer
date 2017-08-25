package main

import (
	"fmt"
	"os"
)

func main() {
	path := validateOsArgs()
	debPrint("CLI path: " + path)
}

func validateOsArgs() string {
	// Expects the following parameters:
	//   1) path_to_dataset
	if len(os.Args) != 2 {
		fmt.Printf("Invalid arguments\nExpected:\n\tpath_to_dataset")
		os.Exit(-1)
	}
	path := os.Args[1]
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Invalid argument 1)\n  Path does not exist: " + path)
		os.Exit(-1)
	}
	return os.Args[1]
}

// DEBUGGING
const (
	CRITICAL = iota
	WARNING
	GENERAL
	VERBOSE
)

var debugLevel = GENERAL

func debPrint(text string) {
	debPrintLevel(text, VERBOSE)
}

func debPrintLevel(text string, level int) {
	switch level {
	case CRITICAL:
		fmt.Printf("[C] " + text)
	case WARNING:
		fmt.Printf("  [W] " + text)
	case GENERAL:
		fmt.Printf("    [G] " + text)
	case VERBOSE:
		fmt.Printf("      [V] " + text)
	}

}
