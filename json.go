package fireball

import (
	"encoding/json"
)

// NewJSONResponse returns a new HTTPResponse in JSON format
func NewJSONResponse(status int, data interface{}, headers map[string]string) (*HTTPResponse, error) {
	if headers == nil {
		headers = JSONHeaders
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	response := NewResponse(status, bytes, headers)
	return response, nil
}

// NewJSONError returns a new HTTPError in JSON format
func NewJSONError(status int, err error, headers map[string]string) (*HTTPError, error) {
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
		HTTPResponse: NewResponse(status, bytes, headers),
		Err:          err,
	}

	return response, nil
}
