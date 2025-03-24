package models

import (
	"fmt"
	"strings"
)

type VersionConstraint struct {
	Operator string `json:"operator"` // e.g., ">=", "<", "==", etc.
	Version  string `json:"version"`  // e.g., "1.2.3"
}

type Dependency struct {
	Name          string               `json:"name"`           // Name of the dependency
	Constraints   *[]VersionConstraint `json:"constraints"`    // Multiple version constraints
	IsUnspecified bool                 `json:"is_unspecified"` // Flag indicating if the version is unspecified
}

func NewDependency(name, constraints string) (Dependency, error) {
	dep := Dependency{Name: name}

	constraints = strings.TrimSpace(constraints)
	if constraints == "" {
		dep.IsUnspecified = true
		return dep, nil
	}

	for _, constraint := range strings.Split(constraints, ",") {
		constraint = strings.TrimSpace(constraint)

		var operator, version string

		switch {
		case strings.HasPrefix(constraint, "=="):
			operator = "exact"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, "=="))
		case strings.HasPrefix(constraint, ">="):
			operator = "gte"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, ">="))
		case strings.HasPrefix(constraint, "<="):
			operator = "lte"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, "<="))
		case strings.HasPrefix(constraint, ">"):
			operator = "gt"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, ">"))
		case strings.HasPrefix(constraint, "<"):
			operator = "lt"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, "<"))
		case strings.HasPrefix(constraint, "~="):
			operator = "wildcard"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, "~="))
		case strings.HasPrefix(constraint, "!="):
			operator = "ne"
			version = strings.TrimSpace(strings.TrimPrefix(constraint, "!="))
		default:
			return Dependency{}, fmt.Errorf("invalid constraint format: %s", constraint)
		}

		if dep.Constraints == nil {
			dep.Constraints = &[]VersionConstraint{}
		}

		*dep.Constraints = append(*dep.Constraints, VersionConstraint{
			Operator: operator,
			Version:  version,
		})
	}

	return dep, nil
}
