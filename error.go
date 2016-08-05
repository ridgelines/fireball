package fireball

import (
	"encoding/json"
)

type HTTPError struct {
	*HTTPResponse
	Err error
}

func NewHTTPError(status int, err error) *HTTPError {
	return &HTTPError{
		HTTPResponse: &HTTPResponse{
			status: status,
		},
		Err: err,
	}
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}

func (e *HTTPError) Body() []byte {
	return []byte(e.Err.Error())
}

type JSONError struct {
	*HTTPError
}

func NewJSONError(status int, err error) *JSONError {
	return &JSONError{
		HTTPError: NewHTTPError(status, err),
	}
}

func (e *JSONError) Body() []byte {
	s := struct {
		Error string
	}{
		Error: e.Err.Error(),
	}

	bytes, err := json.Marshal(s)
	if err != nil {
		return []byte(err.Error())
	}

	return bytes
}
