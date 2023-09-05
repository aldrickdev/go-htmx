package routes

import (
	"github.com/gin-gonic/gin"
)

func Configure(c *gin.Context) {
	gh_pat := c.Query("gh_pat")

	c.SetCookie("github_issues_app", gh_pat)
	c.HTML(200, "index", nil)
}
