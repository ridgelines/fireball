package fireball

import (
	"encoding/json"
)

type HTTPError struct {
	StatusCode int
	Err        error
}

func (e *HTTPError) Error() string {
	return e.Err.Error()
}

func (e *HTTPError) Status() int {
	return e.StatusCode
}

func (e *HTTPError) Response() []byte {
	return []byte(e.Err.Error())
}

// todo: HTMLError

type JSONError struct {
	*HTTPError
}

func (e *JSONError) Response() []byte {
	bytes, err := json.Marshal(e.Err)
	if err != nil {
		return []byte(err.Error())
	}

	// marshal failed, generated `{}`
	if len(bytes) == 2 {
		bytes = e.marshalStruct()
	}

	return bytes
}

func (e *JSONError) marshalStruct() []byte {
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
