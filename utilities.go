package fireball

import (
	"net/http"
)

// Redirect wraps http.Redirect in a ResponseFunc
func Redirect(status int, url string) Response {
	return ResponseFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, status)
	})
}
