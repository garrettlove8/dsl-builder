#ifndef RESULT_H
#define RESULT_H

typedef struct {
	void* data;
	char* errorMsg;
	int errorCode;
	char* successMsg;
} result_t;

result_t* result_new(void*, char*, int, char*);

#endif