package api

import (
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"github.com/aldrickdev/go-htmx/cmd/app/utils"
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

func GetClient() graphql.Client {

	baseClient := http.DefaultClient
	baseClient.Transport = &githubPATRoundTripper{
		next:      http.DefaultTransport,
		githubPAT: utils.Env.GH_PAT,
	}

	return graphql.NewClient("https://api.github.com/graphql", baseClient)
}
