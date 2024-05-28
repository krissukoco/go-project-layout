package http_response

const (
	Code_Unknown  = 99999
	Code_Internal = 90000

	// Authentication

	Code_Unauthenticated = 10001
	Code_Forbidden       = 10003

	// Resources

	Code_InvalidRequest = 20001

	// etc.
)
