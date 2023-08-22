package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/aldrickdev/gin_api_template/envvars"
	"github.com/gin-gonic/gin"
)

var Page *template.Template

const templateLocation = "templates/"
const componentsLocation = "templates/components/"
const indexTemplate = "index.html"

func init() {
	envvars.LoadEnvVars()

	index := fmt.Sprintf("%s%s", templateLocation, indexTemplate)
	name := fmt.Sprintf("%s%s", componentsLocation, "card.html")
	results := fmt.Sprintf("%s%s", componentsLocation, "results.html")

	Page = template.Must(Page.ParseFiles(index, name, results))

	fmt.Println(Page.DefinedTemplates())
}

func main() {
	fmt.Printf("Hello Gin, running in %s environment\n", envvars.Env.ENV)

	if envvars.Env.ENV == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.GET("/", indexPage)
	r.GET("/ping", Ping)
	r.GET("/name", nameComponent)
	r.GET("/issues", getRepoIssues)

	err := r.Run(envvars.Env.PORT)
	if err != nil {
		panic(err)
	}
}

func Ping(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Printf("Body: %s, Error: %v", string(body), err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func getRepoIssues(c *gin.Context) {
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

	err := Page.ExecuteTemplate(c.Writer, indexTemplate, data)
	if err != nil {
		fmt.Print(err)
	}
}

func indexPage(c *gin.Context) {
	err := Page.ExecuteTemplate(c.Writer, indexTemplate, nil)
	if err != nil {
		fmt.Print(err)
	}
}

func nameComponent(c *gin.Context) {
	data := make(map[string]string)
	data["Name"] = "Aldrick"

	temp := "card.html"

	err := Page.ExecuteTemplate(c.Writer, temp, data)
	if err != nil {
		fmt.Print(err)
	}
}
