package lexer_test

import (
	"dsl-builder/internal/lexer"
	"testing"
)

// TestWriteObj ensures successful writing to a statement's in-memory object.
func TestWriteObj(t *testing.T) {
	var want error = nil

	statement, _ := lexer.NewStatement("fakeStatment;")

	testJob := struct {
		Name string
	}{
		Name: "Engineer",
	}

	err := statement.NewObj()
	err = statement.WriteObj(0, testJob)

	if err != want {
		t.Fatalf("%v does not equal %v\n", err, want)
	}
}

// TestWriteObj ensures successful writing to a statement's in-memory object.
func TestNewObj(t *testing.T) {
	var want error = nil

	statement, _ := lexer.NewStatement("fakeStatment;")

	err := statement.NewObj()

	if err != want {
		t.Fatalf("%v does not equal %v\n", err, want)
	}
}
