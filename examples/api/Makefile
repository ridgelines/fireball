.PHONY: deps

deps:
	go install github.com/zpatrick/go-sdata

store:
	go-sdata models/movie.go Movie --package stores > stores/movie_store.go
