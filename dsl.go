package dsl

import (
	"dsl-builder/internal/lexer"
	"fmt"
)

// Run allows a developer to execute a DSL script.
func Run(userScript, sep string) error {
	script := lexer.New(userScript, sep)

	fmt.Println(script)

	return nil
}
