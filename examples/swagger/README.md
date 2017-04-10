# Swagger
This example application adds [Swagger](http://swagger.io/) documentation to the [API example](https://github.com/zpatrick/fireball/tree/master/examples/api). 

## Run Example
From this directory, run:
```
go run main.go
```

By default, if you navigate to `http://localhost:9090/api/`, it will serve Swagger's default [Petstore](http://petstore.swagger.io/) example. 
To use the local configuration, enter `http://localhost:9090/swagger.json` into the **Explore** box on the top right of the page, 
or navigate to `http://localhost:9090/api/?url=http://localhost:9090/swagger.json`
