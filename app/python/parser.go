package python

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"scap/app/models"
	"strings"
)

var skipPrefixes []string = []string{"--hash", "#"}

// shouldSkipLine checks if a line should be skipped based on the provided prefixes.
func shouldSkipLine(line string, prefixes []string) bool {
	if line == "" {
		return true
	}
	for _, prefix := range prefixes {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

func parseRequirementLine(line string) (models.Dependency, error) {
	// Trim initial whitespace
	line = strings.TrimSpace(line)

	// Extract first part before semicolon
	line = strings.SplitN(line, ";", 2)[0]

	// Clean up any remaining whitespace
	line = strings.TrimSpace(line)

	// Regex matching dependency with version constraint
	re := regexp.MustCompile(`^(.*?)\s*(==|>=|<=|>|<|~=|!=)\s*([^#\\;]+)`)

	matches := re.FindStringSubmatch(line)

	if len(matches) == 4 {
		name := strings.TrimSpace(matches[1])
		operator := matches[2]
		version := strings.TrimSpace(matches[3])

		dep, err := models.NewDependency(name, operator+""+version)
		if err != nil {
			return models.Dependency{}, fmt.Errorf("failed to create dependency: %w", err)
		}

		return dep, nil
	}

	// Handle dependency without version condition
	dep, err := models.NewDependency(line, "")
	// dep, err := models.NewDependency(line, "", "")
	if err != nil {
		return models.Dependency{}, fmt.Errorf("failed to create dependency without version: %w", err)
	}
	return dep, nil
}

func ParseRequirementxTxtDependencies(file *os.File) []models.Dependency {
	// Process the file line by line
	var dependencies []models.Dependency
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines or comments or any other prefixes defined above
		if shouldSkipLine(line, skipPrefixes) {
			continue
		}

		// Parse the dependency line properly
		dep, err := parseRequirementLine(line)
		if err != nil {
			fmt.Printf("Skipping line '%s' due to parsing error: %v\n", line, err)
			continue
		}

		dependencies = append(dependencies, dep)
		fmt.Printf("Dependency parsed: %+v\n", dep)
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	return dependencies
}
