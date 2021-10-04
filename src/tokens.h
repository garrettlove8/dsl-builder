#ifndef TOKEN_H
#define TOKEN_H

// token_t represents a single lexical unit (token) in a DSL statement.
typedef struct {
	// lexUnit is a string representation of the token's lexical unit.
	char* lexUnit;

	// fn is the function the carries out the instructions of a speicific
	// token.
	void* (*fn)(void*);

	// statement is the statement of which a token is a part.
	statement_t* statement;

	// previous is the token in a statement directly before the current one.
	token_t* previous;

	// next is the token in a statement directly after the current one.
	token_t* next;
} token_t;

// execute calls the function located at token_t->fn.
// This is how a token's logic is executed by a statement.
void* execute(token_t*);

#endif