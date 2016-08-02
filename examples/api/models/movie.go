package models

type Movie struct {
	ID    string `data:"primary_key"`
	Title string
	Year  int
}
