package token

// token_t represents a single lexical unit (token) in a DSL statement.
type token struct {
	// lexUnit is a string representation of the token's lexical unit.
	lexUnit string

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

// execute calls the function located at token_t->fn.
// This is how a token's logic is executed by a statement.
func execute(token *token) {}
