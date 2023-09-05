package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	ghPat, err := c.Cookie("github_issues_app")

	if err != nil {
		fmt.Println("No Cookie Provided")
		c.Redirect(http.StatusTemporaryRedirect, "/setup")
		return
	}

	fmt.Printf("Cookie exists: %s\n", ghPat)
	c.HTML(200, "index", nil)
}
