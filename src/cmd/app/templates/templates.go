package templates

import (
	"fmt"
	"html/template"

	"github.com/gin-contrib/multitemplate"
)

var ApplicationTemplates *template.Template

const (
	TemplateLocation = "./cmd/app/templates/"

	LayoutLocation     = "layouts/"
	ContentLocation    = "pages/"
	ComponentsLocation = "components/"

	IndexContent     = "index.html"
	ResultsComponent = "results.html"
	IssueComponent   = "issue.html"
)

func LoadTemplates(templateDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	baseLayout := fmt.Sprintf("%s%s", templateDir, "/layouts/base.html")
	content := fmt.Sprintf("%s%s", templateDir, "/content/index.html")
	resultsComponent := fmt.Sprintf("%s%s", templateDir, "/components/issueResults.html")
	issueComponent := fmt.Sprintf("%s%s", templateDir, "/components/issue.html")

	r.AddFromFiles("index", baseLayout, content, resultsComponent, issueComponent)
	r.AddFromFiles("result", resultsComponent, issueComponent)

	return r
}

type Index struct {
	Owner          string
	RepositoryName string
}

type Issue struct {
	Author string
	Labels []string
	Link   string
	Number string
	Title  string
}

type IssueResults struct {
	AfterCursor    string
	BeforeCursor   string
	Issues         []Issue
	Owner          string
	PreviousPage   bool
	RepositoryName string
	RequestMade    bool
	Success        bool
	NextPage       bool
}
