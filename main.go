package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/aldrickdev/gin_api_template/api"
	"github.com/aldrickdev/gin_api_template/envvars"
	"github.com/gin-gonic/gin"
)

var Page *template.Template

const templateLocation = "templates/"
const componentsLocation = "templates/components/"
const indexTemplate = "index.html"

type loggingRoundTripper struct {
	next   http.RoundTripper
	logger io.Writer
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), r.Method, r.URL.String())
	return l.next.RoundTrip(r)
}

type retryRoundTripper struct {
	next       http.RoundTripper
	maxRetries int
	delay      time.Duration
}

func (rr retryRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	var attempts int
	for {
		res, err := rr.next.RoundTrip(r)
		attempts++

		// If we have reached our max retry limit return anyways
		if attempts == rr.maxRetries {
			return res, err
		}

		if err == nil && res.StatusCode < http.StatusInternalServerError {
			return res, err
		}

		select {
		case <-r.Context().Done():
			return res, r.Context().Err()

		case <-time.After(rr.delay):
		}
	}
}

type authRoundTripper struct {
	next             http.RoundTripper
	user, pwd, token string
}

func (a authRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	r.SetBasicAuth(a.user, a.pwd)
	token := fmt.Sprintf("bearer %s", a.token)
	r.Header.Set("Authorization", token)
	return a.next.RoundTrip(r)
}

var httpClient = &http.Client{
	// Transport: &loggingRoundTripper{
	// 	next: http.DefaultTransport,
	// 	logger:              os.Stdout,
	// },
	Transport: &authRoundTripper{
		next: &retryRoundTripper{
			next: &loggingRoundTripper{
				next:   http.DefaultTransport,
				logger: os.Stdout,
			},
			maxRetries: 3,
			delay:      time.Duration(1 * time.Second),
		},
		user:  "bobby",
		pwd:   "pwd",
		token: "tokenbaby",
	},
}

type githubPATRoundTripper struct {
	next      http.RoundTripper
	githubPAT string
}

func (g githubPATRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	token := fmt.Sprintf("bearer %s", g.githubPAT)
	r.Header.Add("Authorization", token)

	return g.next.RoundTrip(r)
}

type RepoIssues struct {
	Repository struct {
		Issues struct {
			Edges []struct {
				Node struct {
					Title     string
					Number    int16
					CreatedAt string
					Author    struct {
						Login string
					}
					Labels struct {
						Nodes []struct {
							Name string
						}
					}
				}
			}
		}
	}
}

func init() {
	envvars.LoadEnvVars()

	index := fmt.Sprintf("%s%s", templateLocation, indexTemplate)
	name := fmt.Sprintf("%s%s", componentsLocation, "card.html")
	results := fmt.Sprintf("%s%s", componentsLocation, "results.html")

	Page = template.Must(Page.ParseFiles(index, name, results))
}

func makeRequest(url string) (int, string) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	res, err := httpClient.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	return res.StatusCode, string(body)
}

func main() {
	// if envvars.Env.ENV == "PROD" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	// r := gin.Default()
	// r.GET("/", indexPage)
	// r.GET("/ping", Ping)
	// r.GET("/name", nameComponent)
	// r.GET("/issues", getRepoIssues)

	// Normal Request
	// code, body := makeRequest("http://httpbin.org/status/200")
	// fmt.Println("\n--- RESPONSE ---\n")
	// fmt.Println("STATUS CODE: ", code)
	// fmt.Println("BODY: ", body)

	// Failing Request, adding retries transport
	// code, body = makeRequest("http://httpbin.org/status/500")
	// fmt.Println("\n--- RESPONSE ---\n")
	// fmt.Println("STATUS CODE: ", code)
	// fmt.Println("BODY: ", body)

	ghPAT := os.Getenv("GH_PAT")
	if ghPAT == "" {
		fmt.Println("No Github PAT provided")
		os.Exit(1)
	}
	fmt.Printf("\"%s\"", ghPAT)

	// baseClient := http.Client{
	// 	Transport: &githubPATRoundTripper{
	// 		next:      http.DefaultTransport,
	// 		githubPAT: ghPAT,
	// 	},
	// }
	baseClient := http.DefaultClient
	baseClient.Transport = &githubPATRoundTripper{
		next:      http.DefaultTransport,
		githubPAT: ghPAT,
	}

	gqlClient := graphql.NewClient("https://api.github.com/graphql", baseClient)
	data, err := api.GetIssues(context.Background(), gqlClient)
	if err != nil {
		panic(err)
	}
	for _, issue := range data.Repository.Issues.Edges {
		fmt.Println(issue.Node.Title)
	}

	// Basic Auth
	// code, body := makeRequest("https://webhook.site/fb39dae1-062d-40e0-ae22-7142f12a7e86")
	// fmt.Println("\n--- RESPONSE ---\n")
	// fmt.Println("STATUS CODE: ", code)
	// fmt.Println("BODY: ", body)

	// gqlClient := graphql.NewClient("https://api.github.com/graphql", http.DefaultClient)
	// data, err := api.GetIssues(context.Background(), gqlClient)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(data.Repository)

	// err = r.Run(envvars.Env.PORT)
	// if err != nil {
	// 	panic(err)
	// }
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
