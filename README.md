# Fireball

<!-- toc -->

- [Overview](#overview)
- [Installation](#installation)
- [Getting Started](#getting-started)
- [Routing](#routing)
  * [Basic](#basic-routing)
  * [Static](#static-routing)
  * [Custom](#custom-routing)
- [Handlers](#handlers)
  * [Basic](#basic-handler)
  * [Error](#error-handler)
  * [Not Found]?
  * [HTML](#html-handler)
  * [JSON](#json-handler)
  * [Authentication](#auth-handler)
  * [Custom](#custom-handler)
- [Views]
  * [Glob Parser]()
  * [Custom]()
- [Testing]
  * [Mock App]()
- [Customization]
  * Not Found
  * Custom Error handling
- [Examples](#examples)
  * [API](/tree/master/examples/api)
  * [Blog](/tree/master/examples/api)


<!-- tocstop -->

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
	app := fireball.NewApp()
	app.Routes = []*fireball.Route{
		&fireball.Route{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": func(c *fireball.Context) (interface{}, error) {
					return "Hello, World!", nil
				},
			},
		},
	}

	fmt.Println("Running on port 8000")
	http.ListenAndServe(":8000", app)
}
```

This will run a new webserver at `localhost:8000`

## Routing
## Static
## Auth
