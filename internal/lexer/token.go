package lexer

// Token represents a single lexical unit (token) in a DSL statement.
type Token struct {
	// Type is the type of token
	Type string

	// Literal is a string representation of the token's lexical unit.
	Literal string
}

// NewToken is a constructor function that returns a pointer to a new toek struct instance.
// Currently, there is not much of a difference between using this function and manually
// creating a new token struct instance. However, this function has been added because
// It is expected that there will be a difference in the future (ex. Validation or some sort).
func NewToken(Type, Literal string) Token {
	// TODO: Instead of taking in a statement it shoudl take in an interface of some kind.

	newToke := Token{
		Type:    Type,
		Literal: Literal,
	}

	return newToke
}

// Execute calls the function located at token_t->fn.
// This is how a token's logic is executed by a statement.
func Execute(token *Token) {}
