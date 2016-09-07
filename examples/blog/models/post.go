package models

import (
	"time"
)

type Post struct {
	Title  string
	Date   time.Time
	Author string
	Body   string
}
