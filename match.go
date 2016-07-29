package fireball

type Match struct {
	Route     *Route
	Handler   Handler
	Variables map[string]string
}
