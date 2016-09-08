# Fireball

[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/zpatrick/fireball/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/zpatrick/fireball)](https://goreportcard.com/report/github.com/zpatrick/fireball)
[![Go Doc](https://godoc.org/github.com/zpatrick/fireball?status.svg)](https://godoc.org/github.com/zpatrick/fireball)


## Overview
Fireball is a micro web framework written in Go. 
Unlike many web frameworks written in Go, handlers in Fireball applications return objects instead instead of writing them directly to a http.ResponseWriter. 
This aims to make handlers feel more like "regular" Go functions and push some of the tediousness of the http layer away from business logic.

## Installation
To install this package, run:
```
go get github.com/zpatrick/fireball
```

## Getting Started
The following snipped shows a simple "Hello, World" application using Fireball:
```
package main

import (
  "github.com/zpatrick/fireball"
  "net/http"
)

func index(c *fireball.Context) (fireball.Response, error) {
  return fireball.NewResponse(200, []byte("Hello, World!"), nil), nil
}

func main() {
  indexRoute := &fireball.Route{
    Path: "/",
    Handlers: map[string]fireball.Handler{
      "GET": index,
    },
  }

  routes := []*fireball.Route{indexRoute}
  app := fireball.NewApp(routes)
  http.ListenAndServe(":8000", app)
}
```

This will run a new webserver at `localhost:8000`

## Handlers
[Handlers](https://godoc.org/github.com/zpatrick/fireball#Handler) perform the business logic associated with requests. 
Handlers take a [Context](https://godoc.org/github.com/zpatrick/fireball#Context) object and returns either a [Response](https://godoc.org/github.com/zpatrick/fireball#Response) or an error.

### HTTP Response
The [HTTP Response](https://godoc.org/github.com/zpatrick/fireball#HTTPResponse) is a simple object that implements the [Response](https://godoc.org/github.com/zpatrick/fireball#Response) interface. 
When the Write call is executed, the specified Body, Status, and Headers will be written to the http.ResponseWriter.

Examples:
```
func Index(c *fireball.Context) (fireball.Response, error) {
    return fireball.NewResponse(200, []byte("Hello, World"), nil), nil
}
```

```
func Index(c *fireball.Context) (fireball.Response, error) {
    html := []byte("<h1>Hello, World</h1>")
    return fireball.NewResponse(200, html, fireball.HTMLHeaders), nil
}
```

### HTTP Error
If a Handler returns a non-nil error, the Fireball Application will call its [ErrorHandler](https://godoc.org/github.com/zpatrick/fireball#App) function. 
By default (if your Application object uses the [DefaultErrorHandler](https://godoc.org/github.com/zpatrick/fireball#DefaultErrorHandler)), the Application will check if the error implements the [Response](https://godoc.org/github.com/zpatrick/fireball#Response) interface. 
If so, the the error's Write function will be called. 
Otherwise, a 500 with the content of err.Error() will be written. 

The [HTTPError](https://godoc.org/github.com/zpatrick/fireball#HTTPError) is a simple object that implements both the [Error](https://golang.org/pkg/builtin/#error) and [Response](https://godoc.org/github.com/zpatrick/fireball#Response) interfaces. 
When the Write is executed, the specified status, error, and headers will be written to the http.ResponseWriter. 

Examples:
```
func Index(c *fireball.Context) (fireball.Response, error) {
    return nil, fmt.Errorf("an error occurred")
}
```
```
func Index(c *fireball.Context) (fireball.Response, error) {
    if err := do(); err != nil {
        return nil, fireball.NewError(500, err, nil)
    }
    
    ...
}
```

## Routing

### Basic Router
By default, Fireball uses the [BasicRouter](https://godoc.org/github.com/zpatrick/fireball#BasicRouter) object to match requests to [Route](https://godoc.org/github.com/zpatrick/fireball#Route) objects.
The Route's Path field determines which URL patterns should be dispached to your Route. 
The Route's Handlers field maps different HTTP methods to different [Handlers](https://godoc.org/github.com/zpatrick/fireball#Handler).

You can use `{variable}` blocks in the Path to match any string that doesn't contain a `"/"` character.
The variables defined in the Route's Path field can be accessed using the [Context](https://godoc.org/github.com/zpatrick/fireball#Context) object.

Example:
```
route := &Fireball.Route{
  Path: "/users/{userID}/orders/{orderID}",
  Methods: map[string]fireball.Handler{
    "GET": printUserOrder,
  },
}

func printUserOrder(c *fireball.Context) (fireball.Response, error) {
    userID := c.PathVariables["userID"]
    orderID := c.PathVariables["orderID"]
    message := fmt.Sprintf("User %s ordered item %s", userID, orderID)
    
    return fireball.NewResponse(200, []byte(message), nil)
}
```

### Static Routing
The built-in [FileServer](https://golang.org/pkg/net/http/#FileServer) can be used to serve static content.
The follow snippet would serve files from the `static` directory:
```
  app := fireball.NewApp(...)
  http.Handle("/", app)

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/static/", http.StripPrefix("/static", fs))
  
  http.ListenAndServe(":8000", nil)
```

If the application workspace contained:
```
app/
    main.go
    static/
        hello_world.txt
```

A request to `/static/hello_world.txt` would serve the desired file.


# HTML Templates
By default, Fireball uses the [GlobParser](https://godoc.org/github.com/zpatrick/fireball#GlobParser) to render HTML templates. 
This object recursively searches a given directory for template files matching the given glob pattern. 
The default root directory is `"views"`, and the default glob pattern is `"*.html"`
The name of the templates are `path/from/root/directory` + `filename`. 

For example, if the filesystem contained:
```
views/
    index.html
    partials/
        login.html
```

The templates names generated would be `"index.html"`, and `"partials/login.html"`.
The [Context](https://godoc.org/github.com/zpatrick/fireball#Context) contains a helper function, [HTML](https://godoc.org/github.com/zpatrick/fireball#Context.HTML), which renders templates as HTML.

Example:
```
func Index(c *fireball.Context) (fireball.Response, error) {
    data := "Hello, World!"
    return c.HTML(200, "index.html", data)
}
```

# Decorators
[Decorators](https://godoc.org/github.com/zpatrick/fireball#Decorator) can be used to wrap additional logic around [Handlers](https://godoc.org/github.com/zpatrick/fireball#Handler). 
Fireball has some built-in decorators:
* [BasicAuthDecorator](https://godoc.org/github.com/zpatrick/fireball#BasicAuthDecorator) adds basic authentication using a specified username and password
* [SessionDecorator](https://godoc.org/github.com/zpatrick/fireball#SessionDecorator) manages [gorilla.Session](http://www.gorillatoolkit.org/pkg/sessions#Session) objects

In addition to Decorators, the [Before](https://godoc.org/github.com/zpatrick/fireball#App) and [After](https://godoc.org/github.com/zpatrick/fireball#App) functions on the [Application](https://godoc.org/github.com/zpatrick/fireball#App) object can be used to perform logic when the request is received and after the response has been sent. 

# Examples
* [JSON](https://github.com/zpatrick/fireball/tree/master/examples/api)
* [Logging](https://github.com/zpatrick/fireball/tree/master/examples/blog)
* [Authentication](https://github.com/zpatrick/fireball/tree/master/examples/blog)
* [HTML Templates](https://github.com/zpatrick/fireball/tree/master/examples/blog)
* [TODO:Redirect](#)
* [TODO: Sessions](#)