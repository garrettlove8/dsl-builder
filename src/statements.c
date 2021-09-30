#include <stdlib.h>
#include <string.h>
#include "statements.h"

#define MODEL_OBJ_CAPACITY 1000

void stfree(statement_t* statement) {
	free(statement->obj);
};

void* stalloc() {
	return calloc(1, MODEL_OBJ_CAPACITY);
};

statement_t NewStatement(char* text) {
	statement_t statement;

	strcpy(statement.text, text);
	statement.obj = stalloc();

	return statement;
};

void DeleteStatement(statement_t* statement) {
	stfree(statement);
};