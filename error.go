package fireball

import (
	"net/http"
)

// HTTPError implements the Response and Error interfaces
type HTTPError struct {
	*HTTPResponse
	Err error
}

// NewError returns a new HTTPError
func NewError(status int, err error, headers map[string]string) *HTTPError {
	return &HTTPError{
		HTTPResponse: NewResponse(status, []byte(err.Error()), headers),
		Err:          err,
	}
}

// Error calls the internal Err.Error function
func (e *HTTPError) Error() string {
	return e.Err.Error()
}

// DefaultErrorHandler is the default ErrorHandler used by an App
// If the error implements the Response interface, it will call its Write function
// Otherwise, a 500 with the error message is returned
func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err, ok := err.(Response); ok {
		err.Write(w, r)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
