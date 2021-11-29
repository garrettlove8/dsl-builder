package lexer

import (
	"fmt"
	"strings"
)

// Lexer maintains the full script and the tokens found within it.
type Lexer struct {
	// Full represents the complete script.
	Full string

	// FullSlice contains the full source code broken up into a slice of characters.
	FullSlice []string

	// Tokens is a slice of the tokens found in the script
	Tokens []Token
}

// NewLexer returns a new Lexer with the provided script assigned to its "Full" field
func NewLexer(script string) Lexer {
	return Lexer{
		Full: script,
	}
}

// Lex reads the script and creates the tokens found in it. The tokens are then assigned
// to the Lexer's Tokens field.
func (l *Lexer) Lex() error {
	l.FullSlice = strings.Split(l.Full, "")
	var nextTokenText string

	for i, v := range l.FullSlice {
		fmt.Printf("lexer:Lex:loop:curChar: %v\n", v)

		if v == " " || v == "\t" || v == "\n" {
			continue
		}

		if s := SymbolTable[nextTokenText]; nextTokenText != "" && s != "" {
			newToken := NewToken(s, nextTokenText)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = ""
		}

		nextTokenText = nextTokenText + v
		nextChar := l.readNext(i)

		if s := l.checkNext(nextChar); s != "" {
			newToken := NewToken("IDENT", nextTokenText)

			l.Tokens = append(l.Tokens, newToken)
			nextTokenText = nextChar
		}
	}

	return nil
}

func (l *Lexer) readNext(i int) string {
	if i+1 < len(l.FullSlice) {
		return l.FullSlice[i+1]
	}

	return ""
}

func (l *Lexer) checkNext(nextChar string) string {
	return SymbolTable[nextChar]
}
