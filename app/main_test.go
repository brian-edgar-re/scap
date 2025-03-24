package main

import (
	"fmt"
	"os"
	"reflect"
	"scap/app/models"
	"scap/app/python"
	"testing"
)

func TestParseRequirementxTxtDependencies(t *testing.T) {
	// Path to the requirements.txt file
	filePath := "../files/python/requirements.txt"

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Parse dependencies using the function from the python package.
	deps := python.ParseRequirementxTxtDependencies(file)

	// Define the expected dependencies.
	expected := []models.Dependency{
		{
			Name: "requests",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "exact",
					Version:  "2.26.0",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "matplotlib",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "exact",
					Version:  "3.6.2",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "about-time",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "exact",
					Version:  "4.2.1",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "ipdb",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "gte",
					Version:  "0.6.0",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "pandas",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "gte",
					Version:  "2.2.0",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "pillow",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "lte",
					Version:  "9.0.0",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "numpy",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "gte",
					Version:  "1.21",
				},
				{
					Operator: "lt",
					Version:  "1.27",
				},
			},
			IsUnspecified: false,
		},
		{
			Name: "joblib",
			Constraints: &[]models.VersionConstraint{
				{
					Operator: "gte",
					Version:  "1.2.0",
				},
				{
					Operator: "lt",
					Version:  "1.4",
				},
			},
			IsUnspecified: false,
		},
		{
			Name:          "xxhash",
			Constraints:   nil,
			IsUnspecified: true,
		},
		{
			Name:          "wurlitzer",
			Constraints:   nil,
			IsUnspecified: true,
		},
		{
			Name:          "setuptools",
			Constraints:   nil,
			IsUnspecified: true,
		},
	}

	// Verify the number of dependencies returned.
	if len(deps) != len(expected) {
		t.Fatalf("expected %d dependencies, got %d", len(expected), len(deps))
	}

	// Compare the entire dependency list using reflect.DeepEqual.
	if !reflect.DeepEqual(deps, expected) {
		t.Errorf("expected dependency list %+v, got %+v", expected, deps)
	}
}
