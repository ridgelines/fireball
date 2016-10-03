package fireball

import (
	"encoding/json"
)

// NewJSONResponse returns a new HTTPResponse in JSON format
func NewJSONResponse(status int, data interface{}) (*HTTPResponse, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response := NewResponse(status, bytes, JSONHeaders)
	return response, nil
}

// NewJSONError returns a new HTTPError in JSON format
func NewJSONError(status int, err error) (*HTTPError, error) {
	e := struct {
		Error string
	}{
		Error: err.Error(),
	}

	bytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}

	response := &HTTPError{
		HTTPResponse: NewResponse(status, bytes, JSONHeaders),
		Err:          err,
	}

	return response, nil
}
