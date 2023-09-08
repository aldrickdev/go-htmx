package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	_, err := c.Cookie("github_issues_app")

	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/setup")
		return
	}

	c.HTML(200, "index", nil)
}
