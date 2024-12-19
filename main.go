package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Path to the requirements.txt file
	filePath := "files/python/requirements.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Process the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// Skip empty lines or comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into name and version
		parts := strings.SplitN(line, "==", 2)
		if len(parts) == 2 {
			name := parts[0]
			version := parts[1]
			fmt.Printf("Name: %s, Version: %s\n", name, version)
		} else {
			// Handle cases where no version is specified
			fmt.Printf("Name: %s, Version: (unspecified)\n", parts[0])
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}
