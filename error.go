package fireball

import (
	"net/http"
)

type HTTPError struct {
	*HTTPResponse
	Err error
}

func NewError(status int, err error, headers map[string]string) *HTTPError {
	return &HTTPError{
		HTTPResponse: NewResponse(status, []byte(err.Error()), headers),
		Err:          err,
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}

func DefaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	if err, ok := err.(Response); ok {
		err.Write(w, r)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
