package routes

import (
	"context"
	"fmt"
	"strings"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/aldrickdev/go-htmx/cmd/app/templates"
	"github.com/aldrickdev/go-htmx/cmd/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

	templateData := templates.IssueResults{
		RequestMade: true,
		Success:     true,
	}

	link := c.Query("link")
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
	//

	owner, repoName, err := extractOwnerRepo(link)
	if err != nil {
		fmt.Println(err)
		templateData.Success = false
		c.HTML(422, "index", templateData)
		return
	}
	templateData.Owner = owner
	templateData.RepositoryName = repoName

	gqlClient := api.GetClient(pat)
	apiData, err := api.GetInitialIssues(context.Background(), gqlClient, 5, owner, repoName)
	if err != nil {
		fmt.Println(err)
		templateData.Success = false
		c.HTML(400, "index", templateData)
		return
	}
	utils.GenerateInitialIssuesTemplateData(&templateData, apiData)

	c.HTML(200, "index", templateData)
}
