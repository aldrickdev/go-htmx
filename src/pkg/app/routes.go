package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Body: %s, Error: %v", string(body), err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func GetRepoIssues(c *gin.Context) {
	repoName := c.Query("repository")

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

	// Get Issue Data
	issue1 := issue{
		Number: "001",
		Title:  "First Issue",
		Author: "Aldrick",
		Labels: []string{
			"Beginner Friendly",
			"Bug",
		},
	}
	issue2 := issue{
		Number: "002",
		Title:  "Second Issue",
		Author: "Aldrick",
		Labels: []string{
			"Fix",
		},
	}

	data := resultsStruct{
		RepositoryName: repoName,
		Issues: []issue{
			issue1, issue2,
		},
	}

	err := ApplicationTemplates.ExecuteTemplate(c.Writer, IndexPage, data)
	if err != nil {
		fmt.Print(err)
	}
}

func Index(c *gin.Context) {
	err := ApplicationTemplates.ExecuteTemplate(c.Writer, IndexPage, nil)
	if err != nil {
		panic(err)
	}
}
