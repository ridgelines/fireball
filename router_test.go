package fireball

import (
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"
)

func nilHandler(*Context) (Response, error) {
	return nil, nil
}

func newRequest(method, path string) *http.Request {
	return &http.Request{
		Method: method,
		URL: &url.URL{
			Path: path,
		},
	}
}

func newRouteWithNilHandler(method, path string) *Route {
	return &Route{
		Path: path,
		Handlers: map[string]Handler{
			method: nilHandler,
		},
	}
}

func TestMethodMatch(t *testing.T) {
	for _, method := range []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"} {
		route := newRouteWithNilHandler(method, "/")
		request := newRequest(method, "/")
		router := NewBasicRouter([]*Route{route})

		match, err := router.Match(request)
		if err != nil {
			t.Fatal(err)
		}

		if match == nil {
			t.Errorf("Error on method '%s': Match was nil", method)
		}
	}
}

func TestPathVariableMatch(t *testing.T) {
	testCases := []struct {
		Route         *Route
		Request       *http.Request
		PathVariables map[string]string
	}{
		{
			Route:         newRouteWithNilHandler("GET", "/"),
			Request:       newRequest("GET", "/"),
			PathVariables: map[string]string{},
		},
		{
			Route:         newRouteWithNilHandler("GET", "/items"),
			Request:       newRequest("GET", "/items"),
			PathVariables: map[string]string{},
		},
		{
			Route:   newRouteWithNilHandler("GET", "/items/:itemID"),
			Request: newRequest("GET", "/items/item34"),
			PathVariables: map[string]string{
				"itemID": "item34",
			},
		},
		{
			Route:   newRouteWithNilHandler("GET", "/items/:itemID/:count"),
			Request: newRequest("GET", "/items/item34/83"),
			PathVariables: map[string]string{
				"itemID": "item34",
				"count":  "83",
			},
		},
	}

	for _, testCase := range testCases {
		router := NewBasicRouter([]*Route{testCase.Route})

		match, err := router.Match(testCase.Request)
		if err != nil {
			t.Fatal(err)
		}

		if v, want := match.PathVariables, testCase.PathVariables; !reflect.DeepEqual(v, want) {
			t.Errorf("\nExpected: %#v \nReceived: %#v", want, v)
		}
	}
}

func TestNilMatch(t *testing.T) {
	testCases := []struct {
		Route   *Route
		Request *http.Request
	}{
		{
			Route:   newRouteWithNilHandler("GET", "/"),
			Request: newRequest("PUT", "/"),
		},
		{
			Route:   newRouteWithNilHandler("GET", "/items"),
			Request: newRequest("GET", "/itemss"),
		},
		{
			Route:   newRouteWithNilHandler("GET", "/items/:itemID"),
			Request: newRequest("GET", "/items/item34/other"),
		},
	}

	for _, testCase := range testCases {
		router := NewBasicRouter([]*Route{testCase.Route})

		match, err := router.Match(testCase.Request)
		if err != nil {
			t.Fatal(err)
		}

		if match != nil {
			t.Errorf("Error on Route '%s': Match was not nil", testCase.Route.Path)
		}
	}
}

func TestConcurrentAccess(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	router := NewBasicRouter([]*Route{
		newRouteWithNilHandler("GET", "/"),
		newRouteWithNilHandler("PUT", "/"),
		newRouteWithNilHandler("POST", "/"),
		newRouteWithNilHandler("DELETE", "/"),
	})

	requests := []*http.Request{
		newRequest("GET", "/"),
		newRequest("PUT", "/"),
		newRequest("POST", "/"),
		newRequest("DELETE", "/"),
	}

	numCalls := 0
	done := make(chan bool)
	for i := 0; i < 5; i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					i := rand.Int() % len(requests)
					req := requests[i]
					router.Match(req)
					numCalls++
				}
			}
		}()
	}

	for numCalls < 1000 {
	}

	close(done)
}
