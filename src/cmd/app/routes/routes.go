package routes

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/aldrickdev/go-htmx/cmd/app/templates"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Body: %s, Error: %v", string(body), err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

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
	owner, repoName, err := extractOwnerRepo(link)
	if err != nil {
		// TODO: generate an error for the html
		fmt.Println(err)
	}

	type issue struct {
		Number string
		Title  string
		Author string
		Labels []string
	}
	type resultsStruct struct {
		RepositoryName string
		Issues         []issue
	}

	gqlClient := api.GetClient()
	apiData, err := api.GetIssues(context.Background(), gqlClient, owner, repoName)
	if err != nil {
		// TODO: generate an error for the html
		fmt.Print("Boy this failed")
		panic(err)
	}
	repository := apiData.GetRepository()
	issues := repository.GetIssues()
	issueEdges := issues.GetEdges()

	templateData := resultsStruct{
		RepositoryName: link,
	}

	for _, edge := range issueEdges {
		node := edge.GetNode()
		Number := fmt.Sprint(node.GetNumber())
		Title := node.GetTitle()
		Author := node.GetAuthor().GetLogin()

		var Labels []string
		for _, l := range node.GetLabels().Nodes {
			Labels = append(Labels, l.GetName())
		}

		templateData.Issues = append(
			templateData.Issues,
			issue{
				Number,
				Title,
				Author,
				Labels,
			},
		)
	}

	err = templates.ApplicationTemplates.ExecuteTemplate(c.Writer, templates.IndexPage, templateData)
	if err != nil {
		fmt.Print(err)
	}
}

func Index(c *gin.Context) {
	err := templates.ApplicationTemplates.ExecuteTemplate(c.Writer, templates.IndexPage, nil)
	if err != nil {
		panic(err)
	}
}
