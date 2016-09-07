package controllers

import (
	"github.com/zpatrick/fireball"
	"github.com/zpatrick/fireball/examples/blog/models"
	"time"
)

type IndexController struct{}

func NewIndexController() *IndexController {
	return &IndexController{}
}

func (h *IndexController) Routes() []*fireball.Route {
	routes := []*fireball.Route{
		{
			Path: "/",
			Handlers: map[string]fireball.Handler{
				"GET": h.index,
			},
		},
	}

	return routes
}

func (h *IndexController) index(c *fireball.Context) (fireball.Response, error) {
	data := struct {
		PinnedPost  *models.Post
		RecentPosts []*models.Post
	}{
		PinnedPost: &models.Post{
			Title:  "Lorem Ipsum",
			Date:   time.Now(),
			Author: "John Doe",
			Body: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Praesent sit amet eleifend sapien. 
			Fusce non facilisis est, in laoreet purus. In mattis urna ut urna interdum ornare. 
			Sed viverra egestas quam, sed porta arcu placerat a. Nullam ut nibh dolor. 
			Integer ultricies id sem sed facilisis. Etiam semper laoreet hendrerit`,
		},
		RecentPosts: []*models.Post{
			{
				Title:  "Duis at suscipit purus",
				Date:   time.Now(),
				Author: "John Doe",
				Body: `Etiam a lacus euismod, pharetra nulla sit amet, condimentum felis. 
				Phasellus varius lectus in ornare vulputate. Integer gravida nisl eget accumsan ullamcorper. 
				Duis efficitur velit nec erat vehicula, vitae bibendum felis eleifend. Nunc cursus, 
				nulla vitae pulvinar pulvinar, nibh lectus scelerisque dui, eu tincidunt lacus tellus sit amet nisi.`,
			},
			{
				Title:  "Etiam Aliquet",
				Date:   time.Now(),
				Author: "John Doe",
				Body: `Nam mollis lacus non lectus vulputate fringilla. Proin quis dui non tortor porta porta at ut urna. 
				Duis eu ex volutpat, dignissim justo eu, molestie ex. Vivamus auctor lorem eu tellus accumsan sollicitudin. 
				Nam enim orci, aliquet sit amet tellus eget, euismod suscipit nunc.`,
			},
			{
				Title:  "Duis id Nibh",
				Date:   time.Now(),
				Author: "John Doe",
				Body: `Morbi interdum tincidunt lorem non consectetur. In hac habitasse platea dictumst. 
				Nunc sodales hendrerit felis id condimentum. Quisque quis interdum lacus. 
				Morbi porta auctor leo, non efficitur felis consequat non.`,
			},
		},
	}

	return c.HTML(200, "index.html", data)
}
