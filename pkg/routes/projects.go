package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/zizouhuweidi/mission-ama/pkg/controller"
)

type (
	projects struct {
		controller.Controller
	}

	projectsTable struct {
		Title string
		Body  string
	}
)

func (c *projects) Get(ctx echo.Context) error {
	page := controller.NewPage(ctx)
	page.Layout = "main"
	page.Name = "projects"
	page.Metatags.Description = "Welcome to the Mission-AMA."
	page.Metatags.Keywords = []string{"AMA", "Missions", "Web"}
	page.Data = c.fetchProjects()

	return c.RenderPage(ctx, page)
}

// fetchPosts is an mock example of fetching posts to illustrate how paging works
func (c *projects) fetchProjects() []post {
	posts := make([]post, 20)

	for k := range posts {
		posts[k] = post{
			Title: fmt.Sprintf("Post example #%d", k+1),
			Body:  fmt.Sprintf("Lorem ipsum example #%d ddolor sit amet, consectetur adipiscing elit. Nam elementum vulputate tristique.", k+1),
		}
	}
	return posts
}
