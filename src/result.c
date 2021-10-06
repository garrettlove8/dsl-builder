#include "result.h"

result_t* result_new(void* data, char* errorMsg, int errorCode, char* successMsg) {
	result_t* result;

	result->data = data;
	result->errorMsg = errorMsg;
	result->errorCode = errorCode;
	result->successMsg = successMsg;

	return result;
};