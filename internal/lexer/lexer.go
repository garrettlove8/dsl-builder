package lexer

import (
	"fmt"
	"strings"
)

// Lexer
type Lexer struct {

	// curPos is the current index value being evaluated in the source code
	// curPos int

	// curPos is the next index value being evaluated in the source code
	// readPos int

	// symbolTable holds all the symbols that are recognized as being a part of the DSL
	SymbolTable SymbolTable

	// Full represents the complete script.
	Full string

	// FullSlice contains the full source code broken up into a slice of characters.
	FullSlice []string

	// Tokens is a slice of the tokens found in the script
	Tokens []Token
}

// NewLexer returns a new Lexer with the provided script assigned to its "Full" field
func NewLexer(st SymbolTable, script string) Lexer {
	return Lexer{
		SymbolTable: st,
		Full:        script,
	}
}

// ParseScript creates Statement instances from the script literal based on the
// script's separator character.
func (l *Lexer) Lex() error {
	fmt.Printf("lexer:Lex:Full: %v\n", l.Full)
	fmt.Printf("lexer:Lex:SymbolTable: %s\n", l.SymbolTable)

	l.FullSlice = strings.Split(l.Full, "")
	var nextTokenText string

	for _, v := range l.FullSlice {
		if v == " " || v == "\t" || v == "\n" {
			continue
		}

		if s := l.SymbolTable.General[nextTokenText]; nextTokenText != "" && s != "" {
			fmt.Println("HERE General")
			newToken := NewToken(s, nextTokenText)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""

			fmt.Printf("lexer:Lex:tokens: %v\n", l.Tokens)
		}

		if s := l.SymbolTable.Identifiers[nextTokenText]; nextTokenText != "" && s != "" {
			fmt.Println("HERE Identifiers")
			newToken := NewToken(s, nextTokenText)

			fmt.Printf("lexer:Lex:newToken: %v\n", newToken)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""

			fmt.Printf("lexer:Lex:tokens: %v\n", l.Tokens)
		}

		if s := l.SymbolTable.Literals[nextTokenText]; nextTokenText != "" && s != "" {
			fmt.Println("HERE Literals")
			newToken := NewToken(s, nextTokenText)

			fmt.Printf("lexer:Lex:newToken: %v\n", newToken)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""

			fmt.Printf("lexer:Lex:tokens: %v\n", l.Tokens)
		}

		if s := l.SymbolTable.Operators[nextTokenText]; nextTokenText != "" && s != "" {
			fmt.Println("HERE Operators")
			newToken := NewToken(s, nextTokenText)

			fmt.Printf("lexer:Lex:newToken: %v\n", newToken)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""

			fmt.Printf("lexer:Lex:tokens: %v\n", l.Tokens)
		}

		if s := l.SymbolTable.Delimiters[nextTokenText]; nextTokenText != "" && s != "" {
			fmt.Println("HERE Delimiters")
			newToken := NewToken(s, nextTokenText)

			fmt.Printf("lexer:Lex:newToken: %v\n", newToken)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""

			fmt.Printf("lexer:Lex:tokens: %v\n", l.Tokens)
		}

		if s := l.SymbolTable.Keywords[nextTokenText]; nextTokenText != "" && s != "" {
			fmt.Println("HERE Keywords")
			newToken := NewToken(s, nextTokenText)

			fmt.Printf("lexer:Lex:newToken: %v\n", newToken)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""

			fmt.Printf("lexer:Lex:tokens: %v\n", l.Tokens)
		}

		nextTokenText = nextTokenText + v

		fmt.Printf("lexer:Lex:nextTokenText: %s\n", nextTokenText)
	}

	return nil
}

func (l *Lexer) getNextPos() {

}

// TODO: This is where I'm at
/*
I have it so the lexer is able to recognize the keywords and delimiters in some scenarios.
However, in this test case:

	CREATE user "Garrett"

It is not able to handle the Garrett" part and just stops lexing. I need to look ahead one
character to check if the next character is in the symbol table, if it is not it should keep
going. On the other hand, if the next character is in the symbol table that means we need to
create a new token with what is currently stored in nextTokenText.

This will hopefully allow us to correctly end up with two additional token in the provided
Test case: {IDENT Garrett} {DQUOTE "}
*/
