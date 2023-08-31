package routes

import (
	"context"
	"fmt"
	"strings"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/gin-gonic/gin"
)

func extractOwnerRepo(input string) (string, string, error) {
	splitLink := strings.Split(input, "/")
	splitLinkLen := len(splitLink)

	if splitLinkLen < 2 || splitLinkLen > 5 {
		return "", "", fmt.Errorf("invalid link provided")
	}

	extracted := splitLink[splitLinkLen-2:]

	return extracted[0], extracted[1], nil
}

func GetRepoIssues(c *gin.Context) {
	// Supported input styles:
	// http://github.com/railwayapp/cli
	// github.com/railwayapp/cli
	// railwayapp/cli

	link := c.Query("link")
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
		RequestMade: true,
		Success:     true,
	}

	owner, repoName, err := extractOwnerRepo(link)
	if err != nil {
		fmt.Println(err)
		templateData.Success = false
		c.HTML(422, "index", templateData)
		return
	}

	gqlClient := api.GetClient()
	apiData, err := api.GetInitialIssues(context.Background(), gqlClient, 5, owner, repoName)
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
	templateData.RepositoryName = repoName
	templateData.BeforeCursor = *apiData.Repository.Issues.PageInfo.GetStartCursor()
	templateData.AfterCursor = *apiData.Repository.Issues.PageInfo.GetEndCursor()
	templateData.NextPage = apiData.Repository.Issues.PageInfo.HasNextPage
	templateData.PreviousPage = apiData.Repository.Issues.PageInfo.HasPreviousPage

	c.HTML(200, "index", templateData)
}
