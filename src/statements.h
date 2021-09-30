#ifndef STATEMENTS_H
#define STATEMENTS_H

typedef struct {
	char text[250];
	void* obj;
} statement_t;

statement_t NewStatement();

void DeleteStatement();

#endif