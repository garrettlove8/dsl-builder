#ifndef RESULT_H
#define RESULT_H

// result_t respresent the result of invoking a function.
typedef struct {
	// data any data returned by the callee function to the caller function.
	void* data;

	// errorMsg is a message describing the error that was encountered.
	// Set to NULL if no error.
	char* errorMsg;

	// errorCode is a reference code for the error.
	// Set to NULL if no error.
	int errorCode;

	// successMsg provides a way to send a custom message to the caller
	// function upon successful invocation.
	char* successMsg;
} result_t;

// result_new constructs a new result.
result_t* result_new(void*, char*, int, char*);

#endif