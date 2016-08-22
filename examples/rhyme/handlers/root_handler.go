package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/zpatrick/fireball"
	"math/rand"
	"net/http"
	"strings"
)

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		&fireball.Route{
			Path: "/",
			Get:  h.Index,
		},
	}

	return routes
}

func (h *RootHandler) Index(c *fireball.Context) (interface{}, error) {
	lines := []string{}
	for i := 0; i < 2; i++ {
		for {
			pair, err := getVerse()
			if err != nil {
				fmt.Println(err)
				continue
			}

			lines = append(lines, pair...)
			break
		}
	}

	context := struct {
		Lines []string
	}{
		Lines: lines,
	}

	return c.HTML(200, "index.html", context)
}

/*
func (h *RootHandler) Index(c *fireball.Context) (interface{}, error) {
	pair1, err := getVerse()
	if err != nil {
		return nil, err
	}

	pair2, err := getVerse()
	if err != nil {
		return nil, err
	}

	lines := []string{
		pair1[0],
		pair1[1],
		pair2[0],
		pair2[1],
	}

	return c.HTML(200, "index.html", lines)
}
*/

func getVerse() ([]string, error) {
	verse1 := getRandomVerse()

	lastWord := getLastWord(verse1)
	rhymes, err := getRhymes(lastWord)
	if err != nil {
		return nil, err
	}

	var verse2 string
	for _, word := range rhymes {
		for i, _ := range rand.Perm(len(Songs)) {
			song := Songs[i]
			for j, _ := range rand.Perm(len(song.Verses)) {
				verse := song.Verses[j]
				if lastWord := getLastWord(verse); strings.ToLower(lastWord) == strings.ToLower(word.Word) {
					verse2 = verse
					break
				}
			}
		}
	}

	if verse2 == "" || verse1 == verse2 {
		return nil, fmt.Errorf("Could not find match for line: '%s' (word '%s'). \nTried to rhyme: %v", verse1, lastWord, rhymes)
	}

	return []string{verse1, verse2}, nil
}

func getLastWord(verse string) string {
	split := strings.Split(verse, " ")
	return split[len(split)-1]
}

func getRandomVerse() string {
	song := Songs[rand.Intn(len(Songs))]
	return song.Verses[rand.Intn(len(song.Verses))]
}

type Word struct {
	Word  string
	Score int
}

func getRhymes(word string) ([]Word, error) {
	url := fmt.Sprintf("http://rhymebrain.com/talk?function=getRhymes&word=%s", word)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	var words []Word
	if err := json.NewDecoder(r.Body).Decode(&words); err != nil {
		return nil, err
	}

	// 300 is perfect match
	for i := 0; i < len(words); i++ {
		if words[i].Score < 250 {
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}

	return words, nil
}
