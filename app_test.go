package fireball

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBeforeIsExecuted(t *testing.T) {
	var executed bool

	app := NewApp(nil)
	app.Before = func(http.ResponseWriter, *http.Request) {
		executed = true
	}

	app.ServeHTTP(httptest.NewRecorder(), newRequest("", ""))
	if !executed {
		t.Fail()
	}
}

func TestAfterIsExecuted(t *testing.T) {
	var executed bool

	app := NewApp(nil)
	app.After = func(http.ResponseWriter, *http.Request) {
		executed = true
	}

	app.ServeHTTP(httptest.NewRecorder(), newRequest("", ""))
	if !executed {
		t.Fail()
	}
}

func TestNotFoundHandlerIsExecuted(t *testing.T) {
	var executed bool

	app := NewApp(nil)
	app.NotFoundHandler = func(http.ResponseWriter, *http.Request) {
		executed = true
	}

	app.ServeHTTP(httptest.NewRecorder(), newRequest("", ""))
	if !executed {
		t.Fail()
	}
}

func TestErrorFromRouterIsHandled(t *testing.T) {
	var executed bool

	app := NewApp(nil)
	app.Router = RouterFunc(func(*http.Request) (*RouteMatch, error) {
		return nil, errors.New("")
	})

	app.ErrorHandler = func(http.ResponseWriter, *http.Request, error) {
		executed = true
	}

	app.ServeHTTP(httptest.NewRecorder(), newRequest("", ""))
	if !executed {
		t.Fail()
	}
}

func TestErrorFromHandlerIsHandled(t *testing.T) {
	var executed bool

	app := NewApp(nil)
	app.Router = RouterFunc(func(*http.Request) (*RouteMatch, error) {
		handler := func(*Context) (Response, error) {
			return nil, errors.New("")
		}

		return &RouteMatch{Handler: handler}, nil
	})

	app.ErrorHandler = func(http.ResponseWriter, *http.Request, error) {
		executed = true
	}

	app.ServeHTTP(httptest.NewRecorder(), newRequest("", ""))
	if !executed {
		t.Fail()
	}
}

func TestResponseWriteIsExecuted(t *testing.T) {
	var executed bool

	app := NewApp(nil)
	app.Router = RouterFunc(func(*http.Request) (*RouteMatch, error) {
		handler := func(*Context) (Response, error) {
			response := ResponseFunc(func(http.ResponseWriter, *http.Request) {
				executed = true
			})

			return response, nil
		}

		return &RouteMatch{Handler: handler}, nil
	})

	app.ServeHTTP(httptest.NewRecorder(), newRequest("", ""))
	if !executed {
		t.Fail()
	}
}
