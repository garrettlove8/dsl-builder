package lexer

import (
	"strings"
)

// script_t represents a complete DSL script and provides an API for running a script.
type Script struct {
	// full represents the complete script.
	Full string

	// sep represents the character on which a developer intends for statements
	// to be terminated.
	Sep string
}

// NewScript is the factory function for creating instances of Script
// from a user's input.
func NewScript(full, sep string) *Script {
	return &Script{
		Full: full,
		Sep:  sep,
	}
}

// ParseScript creates Statement instances from the script literal based on the
// script's separator character.
func (s *Script) ParseScript() ([]*Statement, error) {
	var allStatements []*Statement
	statementStr := strings.Split(s.Full, s.Sep)

	for _, v := range statementStr {
		if v == "" {
			continue
		}

		v = strings.TrimSpace(v)

		newStatement, err := NewStatement(v)
		if err != nil {
			return nil, err
		}

		allStatements = append(allStatements, newStatement)

		// Note: ParseScript does not create any of the buffers in the statement's object slice because
		// The script literal is not being evaluated, only divided up into instance of the Statement
		// struct based on the chosen separator character.
	}

	return allStatements, nil
}
