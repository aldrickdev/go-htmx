package main

import (
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

	if ENV == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	g := gin.Default()
	g.HTMLRender = templates.LoadTemplates()

	g.GET("/", routes.Index)
	g.GET("/setup", routes.Setup)
	g.POST("/setup", routes.Configure)
	g.GET("/ping", routes.Ping)
	g.GET("/issues", routes.GetRepoIssues)
	g.GET("/previous-issues", routes.PreviousIssues)
	g.GET("/next-issues", routes.NextIssues)
	g.GET("/redirect", routes.RedirectToIssue)

	err := g.Run(PORT)
	if err != nil {
		panic(err)
	}
}
