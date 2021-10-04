#ifndef TOKEN_H
#define TOKEN_H


typedef struct {
	char* lexUnit;
	void* (*fn)(void*);
	statement_t* statement;
	token_t* previous;
	token_t* next;
} token_t;

void* execute(token_t*);

#endif