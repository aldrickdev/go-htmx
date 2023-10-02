package templates

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"

	"github.com/gin-contrib/multitemplate"
)

var (
	ApplicationTemplates *template.Template
)

const (
	TemplateLocation = "templates/"

	LayoutLocation     = "layouts/"
	ContentLocation    = "content/"
	ComponentsLocation = "components/"

	BaseLayout            = "base.html"
	IndexContent          = "index.html"
	SetupContent          = "setup.html"
	IssueResultsComponent = "issueResults.html"
	IssueComponent        = "issue.html"
)

func LoadTemplates(filesystem embed.FS) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	baseLayout := fmt.Sprintf("%s%s%s", TemplateLocation, LayoutLocation, BaseLayout)

	indexContent := fmt.Sprintf("%s%s%s", TemplateLocation, ContentLocation, IndexContent)
	setupContent := fmt.Sprintf("%s%s%s", TemplateLocation, ContentLocation, SetupContent)

	resultsComponent := fmt.Sprintf("%s%s%s", TemplateLocation, ComponentsLocation, IssueResultsComponent)
	issueComponent := fmt.Sprintf("%s%s%s", TemplateLocation, ComponentsLocation, IssueComponent)

	r.Add(createTemplate(filesystem, "index", baseLayout, indexContent, resultsComponent, issueComponent))
	r.Add(createTemplate(filesystem, "setup", baseLayout, setupContent))
	r.Add(createTemplate(filesystem, "result", resultsComponent, issueComponent))

	return r
}

func createTemplate(filesystem embed.FS, name string, base string, files ...string) (string, *template.Template) {
	newTemplate := template.New(name)

	data, err := fs.ReadFile(filesystem, base)
	if err != nil {
		log.Fatal("Problem Reading File:", err)
	}

	template.Must(newTemplate.Parse(string(data)))

	if files != nil {
		template.Must(newTemplate.ParseFS(filesystem, files...))
	}

	return name, newTemplate
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
