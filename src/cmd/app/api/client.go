package api

import (
	"fmt"
	"net/http"

	"github.com/Khan/genqlient/graphql"
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

func GetClient(ghPat string) graphql.Client {

	baseClient := http.DefaultClient
	baseClient.Transport = &githubPATRoundTripper{
		next:      http.DefaultTransport,
		githubPAT: ghPat,
	}

	return graphql.NewClient("https://api.github.com/graphql", baseClient)
}
