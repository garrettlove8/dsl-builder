#ifndef STATEMENTS_H
#define STATEMENTS_H

// statement_t is a single statement within a larger DSL script.
typedef struct {
	// full is a pointer to the complete string of text making up the statement.
	char* full;

	// obj is a pointer to the statements memory region.
	// As the statement is executed, this is where the results will be stored.
	void* obj;
} statement_t;

// newStatement constructs a statement object - inluding the allocation of
// its memory region.
statement_t newStatement(char*);

// deleteStatement destructs a statement - including its memory region.
void deleteStatement(statement_t*);

// parseStatement creates the token structures found in the statement base
// on the lexical units.
void parseStatement(statement_t*);

// executeStatement runs the statement's logic by running each tokens function
// keeping track of the effects of each on the statement's memory region.
void executeStatement(statement_t*);

#endif