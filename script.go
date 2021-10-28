package dsl

import (
	"dsl-builder/statement"
	"strings"
)

// script_t represents a complete DSL script and provides an API for running a script.
type script struct {
	// full represents the complete script.
	full string

	// sep represents the character on which a developer intends for statements
	// to be terminated.
	sep string
}

// Run allows a developer to execute a DSL script.
func Run(userScript, sep string) error {
	newScript := script{
		full: userScript,
		sep:  sep,
	}

	_, err := parse(newScript)
	if err != nil {
		return err
	}

	return nil
}

func parse(script script) ([]statement.Statement, error) {
	var allStatements []statement.Statement
	sepScript := strings.Split(script.full, script.sep)

	for _, v := range sepScript {
		trimmedString := strings.TrimSpace(v)
		newStatement, err := statement.NewStatement(trimmedString)
		if err != nil {
			return nil, err
		}
		allStatements = append(allStatements, *newStatement)
	}

	return allStatements, nil
}
