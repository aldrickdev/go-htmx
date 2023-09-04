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
	ContentLocation    = "content/"
	ComponentsLocation = "components/"

	BaseLayout       = "base.html"
	IndexContent     = "index.html"
	SetupContent     = "setup.html"
	ResultsComponent = "results.html"
	IssueComponent   = "issue.html"
)

func LoadTemplates() multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	baseLayout := fmt.Sprintf("%s%s%s", TemplateLocation, LayoutLocation, BaseLayout)

	indexContent := fmt.Sprintf("%s%s%s", TemplateLocation, ContentLocation, IndexContent)
	setupContent := fmt.Sprintf("%s%s%s", TemplateLocation, ContentLocation, SetupContent)

	resultsComponent := fmt.Sprintf("%s%s%s", TemplateLocation, ComponentsLocation, ResultsComponent)
	issueComponent := fmt.Sprintf("%s%s%s", TemplateLocation, ComponentsLocation, IssueComponent)

	r.AddFromFiles("index", baseLayout, indexContent, resultsComponent, issueComponent)
	r.AddFromFiles("setup", baseLayout, setupContent)
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
