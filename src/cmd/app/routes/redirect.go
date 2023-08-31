package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RedirectToIssue(c *gin.Context) {
	owner := c.Query("owner")
	repoName := c.Query("repoName")
	issueNum := c.Query("issueNum")

	link := fmt.Sprintf("https://github.com/%s/%s/issues/%s", owner, repoName, issueNum)

	c.Header("HX-Redirect", link)
}
