package main

import (
	"fmt"
	"html/template"

	"github.com/aldrickdev/go-htmx/cmd/app/routes"
	"github.com/aldrickdev/go-htmx/cmd/app/templates"
	"github.com/aldrickdev/go-htmx/cmd/app/utils"
	"github.com/gin-gonic/gin"
)

var (
	ENV    string
	PORT   string
	GH_PAT string
)

func init() {
	utils.LoadEnvVars()

	ENV = utils.Env.ENV
	PORT = utils.Env.PORT
	GH_PAT = utils.Env.GH_PAT
	if GH_PAT == "" {
		panic("No Github PAT provided")
	}
}

func main() {
	g := gin.Default()
	templates.ApplicationTemplates = template.Must(
		template.ParseFiles(
			fmt.Sprintf("%s%s", templates.PagesLocation, templates.IndexPage),
			fmt.Sprintf("%s%s", templates.ComponentsLocation, templates.ResultsComponent),
		),
	)

	g.GET("/", routes.Index)
	g.GET("/ping", routes.Ping)
	g.GET("/issues", routes.GetRepoIssues)

	err := g.Run(PORT)
	if err != nil {
		panic(err)
	}
}
