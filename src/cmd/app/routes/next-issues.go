package routes

import (
	"context"
	"fmt"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/aldrickdev/go-htmx/cmd/app/templates"
	"github.com/aldrickdev/go-htmx/cmd/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func NextIssues(c *gin.Context) {
	owner := c.Query("owner")
	repoName := c.Query("repoName")
	cursor := c.Query("cursor")

	templateData := templates.IssueResults{
		Owner:          owner,
		RepositoryName: repoName,
		RequestMade:    true,
		Success:        true,
	}

	ghpat, err := c.Cookie("github_issues_app")
	if err != nil {
		fmt.Println(err)
		templateData.Success = false
		c.HTML(422, "index", templateData)
		return
	}

	// Abstract
	token, err := jwt.Parse(ghpat, func(token *jwt.Token) (interface{}, error) {
		return utils.Env.JWT_SIGNING_KEY, nil
	})
	if err != nil {
		fmt.Println("Error parsing token: ", err)
		templateData.Success = false
		c.HTML(422, "index", templateData)
		return
	}

	pat, err := token.Claims.GetSubject()
	if err != nil {
		fmt.Println("Error trying to get subject: ", err)
		templateData.Success = false
		c.HTML(422, "index", templateData)
		return
	}

	gqlClient := api.GetClient(pat)
	apiData, err := api.GetNextIssues(context.Background(), gqlClient, 5, owner, repoName, cursor)
	if err != nil {
		fmt.Println(err)
		templateData.Success = false
		c.HTML(400, "index", templateData)
		return
	}
	utils.GenerateNextIssuesTemplateData(&templateData, apiData)

	c.HTML(200, "result", templateData)
}
