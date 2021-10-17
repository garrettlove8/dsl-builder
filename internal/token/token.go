package token

import "dsl-builder/internal/statement"

// token_t represents a single lexical unit (token) in a DSL statement.
type token struct {
	// literal is a string representation of the token's lexical unit.
	literal string

	// fn is the function the carries out the instructions of a speicific
	// token.
	fn interface{} // TODO: this should be changed to not be an interface type

	// statement is the statement of which a token is a part.
	statement interface{} // TODO: this should be changed to not be an interface type

	// previous is the token in a statement directly before the current one.
	previous *token

	// next is the token in a statement directly after the current one.
	next *token
}

// NewToken is a constructor function that returns a pointer to a new toek struct instance.
// Currently, there is not much of a difference between using this function and manually
// creating a new token struct instance. However, this function has been added because
// It is expected that there will be a difference in the future (ex. Validation or some sort).
func NewToken(literal string, fn func(...interface{}) interface{}, statement *statement.Statement, previous *token, next *token) (*token, error) {
	// TODO: Instead of taking in a statement it shoudl take in an interface of some kind.

	newToke := token{
		literal:   literal,
		fn:        fn,
		statement: statement,
		previous:  previous,
		next:      next,
	}

	return &newToke, nil
}

// Execute calls the function located at token_t->fn.
// This is how a token's logic is executed by a statement.
func Execute(token *token) {}
