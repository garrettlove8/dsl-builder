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

statement_t newStatement(char* text) {
	statement_t statement;

	// strcpy(statement.full, text);
	statement.obj = stalloc();

	return statement;
};

void deleteStatement(statement_t* statement) {
	stfree(statement);
};