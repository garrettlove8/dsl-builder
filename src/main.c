#include <stdio.h>
#include <stdlib.h>
#include "tokens.h"
#include "statements.h"

int main(void) {
	statement_t statement = NewStatement("CREATE job Manager");
	printf("statement: %s\n", statement.text);

	for(int n = 0; n < 1000; ++n) {
		printf("statement[%d]: %d\n", n, ((char *)statement.obj)[n]);
	}

	DeleteStatement(&statement);

	return EXIT_SUCCESS;
}