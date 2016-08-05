package fireball

type Status interface {
	Status() int
}

type Body interface {
	Body() []byte
}

type Headers interface {
	Headers() map[string]string
}

type HTTPResponse struct {
	status  int
	body    []byte
	headers map[string]string
}

func NewHTTPResponse(status int, body []byte, headers map[string]string) *HTTPResponse {
	if headers == nil {
		headers = map[string]string{}
	}

	return &HTTPResponse{
		status:  status,
		body:    body,
		headers: headers,
	}
}

func (r *HTTPResponse) Status() int {
	return r.status
}

func (r *HTTPResponse) Body() []byte {
	return r.body
}

func (r *HTTPResponse) Headers() map[string]string {
	return r.headers
}
