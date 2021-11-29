package lexer_test

import (
	"dsl-builder/internal/lexer"
	"testing"
)

// TestWriteObj ensures successful writing to a statement's in-memory object.
func TestLexer(t *testing.T) {
	lexer.AddSymbol("CREATE", "CREATE")
	lexer.AddSymbol("user", "user")
	lexer.AddSymbol("DQUOTE", `"`)
	lexer.AddSymbol("IDENT", "IDENT")

	lex := lexer.NewLexer(`
		CREATE user "Garrett"
	`)

	err := lex.Lex()

	if err != nil {
		t.Fatalf("%v does not equal nil\n", err)
	}
}
