package main

import (
	"fmt"
	"os"
	"scap/app/python"
)

func main() {
	// Path to the requirements.txt file
	filePath := "../files/python/requirements.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Extract dependencies from file
	dependencies := python.ParseRequirementxTxtDependencies(file)
	fmt.Println(dependencies)

}
