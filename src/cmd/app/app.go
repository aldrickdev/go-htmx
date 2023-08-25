package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/aldrickdev/go-htmx/pkg/app"
	"github.com/aldrickdev/go-htmx/pkg/utils"
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

	app.ApplicationTemplates = template.Must(
		template.ParseFiles(
			fmt.Sprintf("%s%s", app.PagesLocation, app.IndexPage),
			fmt.Sprintf("%s%s", app.ComponentsLocation, app.ResultsComponent),
		),
	)

	fmt.Println(app.ApplicationTemplates.DefinedTemplates())

	g.GET("/", app.Index)
	g.GET("/ping", app.Ping)
	g.GET("/issues", app.GetRepoIssues)

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
