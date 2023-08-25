package main

import (
	"fmt"
	"html/template"
	"net/http"

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

type githubPATRoundTripper struct {
	next      http.RoundTripper
	githubPAT string
}

func (g githubPATRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	token := fmt.Sprintf("bearer %s", g.githubPAT)
	r.Header.Add("Authorization", token)

	return g.next.RoundTrip(r)
}

func init() {
	utils.LoadEnvVars()
	ENV = utils.Env.ENV
	PORT = utils.Env.PORT
	GH_PAT = utils.Env.GH_PAT
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
	// ghPAT := os.Getenv("GH_PAT")
	// if ghPAT == "" {
	// 	fmt.Println("No Github PAT provided")
	// 	os.Exit(1)
	// }
	// fmt.Printf("\"%s\"", ghPAT)

	// baseClient := http.DefaultClient
	// baseClient.Transport = &githubPATRoundTripper{
	// 	next:      http.DefaultTransport,
	// 	githubPAT: ghPAT,
	// }

	// gqlClient := graphql.NewClient("https://api.github.com/graphql", baseClient)
	// data, err := api.GetIssues(context.Background(), gqlClient)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, issue := range data.Repository.Issues.Edges {
	// 	fmt.Println(issue.Node.Title)
	// }

	// err = r.Run(envvars.Env.PORT)
	// if err != nil {
	// 	panic(err)
	// }
}
