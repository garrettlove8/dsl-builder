package lexer_test

import (
	"dsl-builder/internal/lexer"
	"testing"
)

// TestWriteObj ensures successful writing to a statement's in-memory object.
func TestLexer(t *testing.T) {
	st := lexer.NewSymbolTable()
	st.AddSymbol("Keywords", "CREATE", "CREATE")
	st.AddSymbol("Keywords", "user", "user")
	st.AddSymbol("Delimiters", "DQUOTE", `"`)

	lex := lexer.NewLexer(st, `
		CREATE user "Garrett"
	`)

	err := lex.Lex()

	if err != nil {
		t.Fatalf("%v does not equal nil\n", err)
	}
}
