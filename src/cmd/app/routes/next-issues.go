package routes

import (
	"context"
	"fmt"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/aldrickdev/go-htmx/cmd/app/templates"
	"github.com/aldrickdev/go-htmx/cmd/app/utils"
	"github.com/gin-gonic/gin"
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

	gqlClient := api.GetClient(utils.Env.GH_PAT)
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
