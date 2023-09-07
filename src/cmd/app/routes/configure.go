package routes

import (
	"github.com/gin-gonic/gin"
)

func Configure(c *gin.Context) {
	ghpat, found := c.GetPostForm("ghpat")
	if !found {
		c.HTML(200, "setup", nil)
		return
	}

	c.SetCookie("github_issues_app", ghpat, 0, "/", "", true, true)
	c.Header("HX-Redirect", "/")
}
