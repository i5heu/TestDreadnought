package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/i5heu/TestDreadnought/internal/orchestrator"
)

func main() {
	fmt.Println("TestDreadnought", "v0.1.1")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <test-root-directory> <optional: subset path relative to config-directory>")
		return
	}

	// Get the config directory path and optional subset path
	configDir := os.Args[1]
	var subSet string
	if len(os.Args) >= 3 {
		subSet = os.Args[2]
	}

	// Check if subSet exists inside configDir
	if subSet != "" {
		subSetPath := filepath.Join(configDir, subSet)
		if _, err := os.Stat(subSetPath); os.IsNotExist(err) {
			fmt.Printf("Subset directory %s does not exist inside %s\n", subSet, configDir)
			return
		}
	}

	// Run the tests
	err := orchestrator.RunTests(configDir, subSet)
	if err != nil {
		fmt.Println(err)
	}
}
