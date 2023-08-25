package app

import (
	"html/template"
)

var ApplicationTemplates *template.Template

const (
	PagesLocation      = "templates/pages/"
	ComponentsLocation = "templates/components/"
	IndexPage          = "index.html"
	ResultsComponent   = "results.html"
)
