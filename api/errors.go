package api

type AuthorizationError struct {
	StatusCode int
	Message    string
}

func (e *AuthorizationError) Error() string {
	return e.Message
}

type BMCAPIError struct {
	StatusCode int
	Message    string
}

func (e *BMCAPIError) Error() string {
	return e.Message
}

type NoSessionError struct {
	Message string
}

func (e *NoSessionError) Error() string {
	return "no session found"
}
