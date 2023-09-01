package routes

import (
	"context"
	"fmt"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/gin-gonic/gin"
)

func PreviousIssues(c *gin.Context) {
	owner := c.Query("owner")
	repoName := c.Query("repoName")
	cursor := c.Query("cursor")

	type issue struct {
		Number string
		Title  string
		Author string
		Link   string
		Labels []string
	}
	type resultsStruct struct {
		Owner          string
		RepositoryName string
		Issues         []issue
		RequestMade    bool
		Success        bool
		BeforeCursor   string
		AfterCursor    string
		NextPage       bool
		PreviousPage   bool
	}
	templateData := resultsStruct{
		RepositoryName: repoName,
		RequestMade:    true,
		Success:        true,
	}

	gqlClient := api.GetClient()
	apiData, err := api.GetPreviousIssues(context.Background(), gqlClient, 5, owner, repoName, cursor)
	if err != nil {
		fmt.Println(err)
		templateData.Success = false
		c.HTML(400, "index", templateData)
		return
	}
	repository := apiData.GetRepository()
	issues := repository.GetIssues()
	issueEdges := issues.GetEdges()

	for _, edge := range issueEdges {
		node := edge.GetNode()
		Number := fmt.Sprint(node.GetNumber())
		Title := node.GetTitle()
		AuthorStruct := *node.GetAuthor()
		AuthorName := AuthorStruct.GetLogin()
		Link := fmt.Sprintf("?owner=%s&repoName=%s&issueNum=%s", owner, repoName, Number)

		var Labels []string
		for _, l := range node.GetLabels().Nodes {
			Labels = append(Labels, l.GetName())
		}

		templateData.Issues = append(
			templateData.Issues,
			issue{
				Number,
				Title,
				AuthorName,
				Link,
				Labels,
			},
		)
	}

	templateData.Owner = owner
	templateData.BeforeCursor = *apiData.Repository.Issues.PageInfo.GetStartCursor()
	templateData.AfterCursor = *apiData.Repository.Issues.PageInfo.GetEndCursor()
	templateData.NextPage = apiData.Repository.Issues.PageInfo.GetHasNextPage()
	templateData.PreviousPage = apiData.Repository.Issues.PageInfo.GetHasPreviousPage()

	c.HTML(200, "result", templateData)
}
