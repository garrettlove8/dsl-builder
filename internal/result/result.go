package result

// result respresent the result of invoking a function.
type result struct {
	// data any data returned by the callee function to the caller function.
	data interface{}

	// errorMsg is a message describing the error that was encountered.
	// Set to NULL if no error.
	errorMsg string

	// errorCode is a reference code for the error.
	// Set to NULL if no error.
	errorCode int

	// successMsg provides a way to send a custom message to the caller
	// function upon successful invocation.
	successMsg string
}

// result_new constructs a new result.
func (r *result) result_new(data interface{}, errorMsg string, errorCode int, successMsg string) (*result, error) {
	result := result{}

	result.data = data
	result.errorMsg = errorMsg
	result.errorCode = errorCode
	result.successMsg = successMsg

	return &result, nil
}
