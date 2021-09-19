#include <stdio.h>
#include <stdlib.h>
#include <stdarg.h>
#include "token.h"

/*
	Challenges:
		- Users have created a semantic model, we need to be able to use their
		model to run the DSL	
	
	token_t -> userObj
	Maybe not: Use polymorphism to allow the caller to override the function I am defining as the action function of the token type
*/

// --- Job Model ---
typedef struct {
	char* name;
} job_t;

int job_create(void* objPtr, ...);

int main(void) {
	symbols_t symbolTable;
	symbolTable.funcOne = *job_create;

	token_t myStruct;
	job_t myJob;

	// symbolTable.funcOne(&myJob, "hey1", "hey2", "hey3", "hey4", "hey5", "hey6", "hey7", "hey8", NULL);
	// symbolTable.funcOne(&myJob, "hey1", "hey2");
	symbolTable.funcOne(&myJob, "Manager", NULL);
	// symbolTable.funcOne(&myJob, NULL);
	// job_create(&myJob, "Garrett");

	printf("Job name: %s", myJob.name);

	return EXIT_SUCCESS;
}

int job_create(void* objPtr, ...) {
	char* name;
	job_t* jobPtr;

	jobPtr = (job_t*)objPtr;

	int i = 0;

	va_list args;
	printf("args bytes: %lu\n", sizeof(args));
	
	va_start(args, objPtr);
	void* arg;

	while(arg) {
		arg = va_arg(args, void*);
		if (arg == NULL) {
			break;
		}

		printf("arg: %s\n", (char*)arg);

		if (i == 0) {
			jobPtr->name = (char*)arg;
		}

		i++;
	}

	va_end(args);

	return EXIT_SUCCESS;
}