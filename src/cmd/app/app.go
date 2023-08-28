package main

import (
	"fmt"

	"github.com/aldrickdev/go-htmx/cmd/app/routes"
	"github.com/aldrickdev/go-htmx/cmd/app/utils"
	"github.com/gin-contrib/multitemplate"
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

	if ENV == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
}

func main() {
	g := gin.Default()
	g.HTMLRender = loadTemplates("./templates/new")

	g.GET("/", routes.Index)
	g.GET("/ping", routes.Ping)
	g.GET("/issues", routes.GetRepoIssues)

	err := g.Run(PORT)
	if err != nil {
		panic(err)
	}
}

func loadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	baseLayout := fmt.Sprintf("%s%s", templateDir, "/layouts/base.html")
	content := fmt.Sprintf("%s%s", templateDir, "/content/index.html")
	component := fmt.Sprintf("%s%s", templateDir, "/components/issueResults.html")

	r.AddFromFiles("index", baseLayout, content, component)

	return r
}
