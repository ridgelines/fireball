# Fireball

[![Build Status](https://travis-ci.org/urfave/cli.svg?branch=master)](https://travis-ci.org/urfave/cli)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/fireball/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/fireball)](https://goreportcard.com/report/github.com/zpatrick/fireball)
[![Go Doc](https://godoc.org/github.com/zpatrick/fireball?status.svg)](https://godoc.org/github.com/zpatrick/fireball)


## Overview
A micro web framework written in Go

## Installation
To install this package, run:
```
go get github.com/zpatrick/fireball
```

## Getting Started
The following example shows how to write a simple "Hello, World" application using Fireball. 
To run this example, create a `main.go` file with the following:
```
package main

import (
  "fmt"
  "github.com/zpatrick/fireball"
  "net/http"
)

func main() {
  routes := []*fireball.Route{
    &fireball.Route{
      Path: "/",
      Handlers: map[string]fireball.Handler{
        "GET": func(c *fireball.Context) (interface{}, error) {
          return "Hello, World!", nil
        },
      },
    },
  }

  app := fireball.NewApp(routes)

  fmt.Println("Running on port 8000")
  http.ListenAndServe(":8000", app)
}
```

This will run a new webserver at `localhost:8000`

## Routing

### Basic Routes
Fireball uses **Route** objects to dispatch requests to handlers.
The **Path** field determines which URLs should be dispached to your Route. 
You can use `{variable}` blocks in the Path to match any string that doesn't contain a `"/"` character.
Routes also contain a maping of http methods to different **Handlers** which perform the business logic.

For example, take the following Fireball Routes:
```
routes := []&Fireball.Route{
    &Fireball.Route{
        Path: "/movies/{movie}",
        Methods: map[string]fireball.Handler{
            "GET": getMovie,
            "POST": createMovie,
        }
    },
    &Fireball.Route{
        Path: "/users/{user}/orders/{order}",
        Methods: map[string]fireball.Handler{
            "GET": getUserOrder,
        }
    },
}
```

The following requests would be routed as so:

| URL | Method | Action | Path Variables |
| :-: | :-: | :-: | :-: |
| `/movies/m1` | GET | `getMovie()` is called | `movie="m1"` |
| `/movies/myMovie` | POST | `createMovie()` is called | `id="myMovie"` |
| `/movies/m1/view` | GET | no action | n/a |
| `/movies` | GET | no action | n/a | 
| `/movies/m1` | DELETE | no action | n/a |
| `/users/u1/orders/o1` | GET | `getUserOrder()` is called| `user="u1"`, `order="o1"` |

For more information on Handlers (including how to access the variable(s) specified in the URL), see the [Handlers](#Handlers) section.


### Static Routing
Static content routing can be accomplished with the built-in **http.FileServer**.
The follow snippet would serve files from the `static` directory (which must be in the same working directory as the snippet).
```
  app := fireball.NewApp(...)
  http.Handle("/", app)

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static", fs))
  
  http.ListenAndServe(":8000", nil)
```

For example, if your application workspace contained:
```
app/
    main.go
    static/
        hello_world.txt
```

Making a request to `/static/hello_world.txt` would serve the proper file for you. 

### Custom Routing
You can implement a custom **Router** for your Fireball applications as long as it satisfies the proper interface:
```
type Router interface {
  Match(*http.Request) (*RouteMatch, error)
}
```

The `Router.Match()` function is called each time an incomping request is made.
If it matches the request to a Route, it should  should return a [RouteMatch](https://godoc.org/github.com/zpatrick/fireball#RouteMatch) object. 
Otherwise, it should return `nil` (in which case a `NotFound` response is sent) 

```
app := fireball.NewApp(nil)
app.Router = MyCustomRouter
...
```
## Handlers
[Handlers](https://godoc.org/github.com/zpatrick/fireball#Handler) perform the business logic associated with requests.
Unlike most web frameworks written in Go, Handlers in Fireball applications return responses and errors instead of writing them directly to a `http.ResponseWriter`. 
This aims to make Handlers feel more like "regular" Go functions, and it pushes some of the tediousness of the http layer away from the business logic. 

All Handlers take a [Context](https://godoc.org/github.com/zpatrick/fireball#Context) object. 
This object provides some helper functions along with access to the originating `http.Request`. 
Handlers must return a response (an `interface{}`) or an `error`. 
Fireball will attempt to return an appropriate response based on what was returned from the Handler. 
It is recommended (but not required) that Handlers return [Response](#Response) objects where possible.

### Response Objects
A Handler can return any type of object as a response. 
```
func Index(c *fireball.Context) (interface{}, error) {
    return "Hello, World!", nil
}
```

By default, response objects will be rendered in default string format, contain a 200 status code, and not contain headers. This behavior can be changed depending on the type of response object that is returned:
* The rendered response body can be overwritten by implementing the  [Body](https://godoc.org/github.com/zpatrick/fireball#Body) interface
* The status code can be overwritten by implementing the [Status](https://godoc.org/github.com/zpatrick/fireball#Status) interface
* The response headers can be overwritten by implementing the  [Headers](https://godoc.org/github.com/zpatrick/fireball#Headers) interface

Fireball has a built-in [HTTPResponse](https://godoc.org/github.com/zpatrick/fireball#HTTPResponse) object that fulfills all of these interfaces. 
```
func Index(c *fireball.Context) (interface{}, error) {
    response := fireball.NewHTTPResponse(200, []byte("Hello, World"), nil)
    return response, nil
}
```

### Error Objects
If a Handler returns a non-nil error, Fireball will use error as a response instead of the response object. 
```
func Index(c *fireball.Context) (interface{}, error) {
    return nil, fmt.Errorf("an error occurred")
}
```

By default, error objects will be rendered using the `Error()` function, contain a 500 status code, and not contain headers. The same rules that apply to response objects apply to errors: the behavior can be changed depending on the type of error object that is returned:
* The rendered response body can be overwritten by implementing the  [Body](https://godoc.org/github.com/zpatrick/fireball#Body) interface
* The status code can be overwritten by implementing the [Status](https://godoc.org/github.com/zpatrick/fireball#Status) interface
* The response headers can be overwritten by implementing the  [Headers](https://godoc.org/github.com/zpatrick/fireball#Headers) interface

Fireball has a built-in [HTTPError](https://godoc.org/github.com/zpatrick/fireball#HTTPError) object that fulfills all of these interfaces. 
```
func Index(c *fireball.Context) (interface{}, error) {
    err := fireball.NewHTTPError(500, []byte("an error occurred"), nil)
    return nil, err
}
```

### JSON
Fireball has built-in support to send responses in JSON format:

```
func Index(c *fireball.Context) (interface{}, error) {
    data := map[string]string{"hello", "world"}
    return fireball.NewJSONResponse(200, data, nil)
}
```

Along with sending errors in JSON format:
```
func Index(c *fireball.Context) (interface{}, error) {
    err := fmt.Errorf("an error occurred")
    return nil, fireball.NewJSONError(500, err, nil)
}
```

### HTML
The [Context.HTML](https://godoc.org/github.com/zpatrick/fireball#Context.HTML) function can be used to render HTML from templates in your application. 
Please see the [Templates](#templates) section for information on how to set that up.

This function takes a status code, name of the template file, and data (an `interface{}`) to send to the template renderer. 
```
  return c.HTML(200, "index.html", data)
```

### Authentication
### Custom


## Template Parsing


