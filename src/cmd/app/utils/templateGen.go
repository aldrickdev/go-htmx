package utils

import (
	"fmt"

	"github.com/aldrickdev/go-htmx/cmd/app/api"
	"github.com/aldrickdev/go-htmx/cmd/app/templates"
)

func GenerateInitialIssuesTemplateData(templateData *templates.IssueResults, apiData *api.GetInitialIssuesResponse) {
	repository := apiData.GetRepository()
	issues := repository.GetIssues()
	issueEdges := issues.GetEdges()

	if len(issueEdges) == 0 {
		templateData.Issues = nil

		templateData.BeforeCursor = ""
		templateData.AfterCursor = ""
		templateData.NextPage = false
		templateData.PreviousPage = false

		return
	}

	for _, edge := range issueEdges {
		node := edge.GetNode()
		Number := fmt.Sprint(node.GetNumber())
		Title := node.GetTitle()
		AuthorStruct := *node.GetAuthor()
		Author := AuthorStruct.GetLogin()
		Link := fmt.Sprintf("?owner=%s&repoName=%s&issueNum=%s", templateData.Owner, templateData.RepositoryName, Number)

		var Labels []string
		for _, l := range node.GetLabels().Nodes {
			Labels = append(Labels, l.GetName())
		}

		templateData.Issues = append(
			templateData.Issues,
			templates.Issue{
				Author: Author,
				Labels: Labels,
				Link:   Link,
				Number: Number,
				Title:  Title,
			},
		)
	}

	templateData.BeforeCursor = *apiData.Repository.Issues.PageInfo.GetStartCursor()
	templateData.AfterCursor = *apiData.Repository.Issues.PageInfo.GetEndCursor()
	templateData.NextPage = apiData.Repository.Issues.PageInfo.GetHasNextPage()
	templateData.PreviousPage = apiData.Repository.Issues.PageInfo.GetHasPreviousPage()
}

func GenerateNextIssuesTemplateData(templateData *templates.IssueResults, apiData *api.GetNextIssuesResponse) {
	repository := apiData.GetRepository()
	issues := repository.GetIssues()
	issueEdges := issues.GetEdges()

	for _, edge := range issueEdges {
		node := edge.GetNode()
		Number := fmt.Sprint(node.GetNumber())
		Title := node.GetTitle()
		AuthorStruct := *node.GetAuthor()
		Author := AuthorStruct.GetLogin()
		Link := fmt.Sprintf("?owner=%s&repoName=%s&issueNum=%s", templateData.Owner, templateData.RepositoryName, Number)

		var Labels []string
		for _, l := range node.GetLabels().Nodes {
			Labels = append(Labels, l.GetName())
		}

		templateData.Issues = append(
			templateData.Issues,
			templates.Issue{
				Author: Author,
				Labels: Labels,
				Link:   Link,
				Number: Number,
				Title:  Title,
			},
		)
	}

	templateData.BeforeCursor = *apiData.Repository.Issues.PageInfo.GetStartCursor()
	templateData.AfterCursor = *apiData.Repository.Issues.PageInfo.GetEndCursor()
	templateData.NextPage = apiData.Repository.Issues.PageInfo.GetHasNextPage()
	templateData.PreviousPage = apiData.Repository.Issues.PageInfo.GetHasPreviousPage()
}

func GeneratePreviousIssuesTemplateData(templateData *templates.IssueResults, apiData *api.GetPreviousIssuesResponse) {
	repository := apiData.GetRepository()
	issues := repository.GetIssues()
	issueEdges := issues.GetEdges()

	for _, edge := range issueEdges {
		node := edge.GetNode()
		Number := fmt.Sprint(node.GetNumber())
		Title := node.GetTitle()
		AuthorStruct := *node.GetAuthor()
		Author := AuthorStruct.GetLogin()
		Link := fmt.Sprintf("?owner=%s&repoName=%s&issueNum=%s", templateData.Owner, templateData.RepositoryName, Number)

		var Labels []string
		for _, l := range node.GetLabels().Nodes {
			Labels = append(Labels, l.GetName())
		}

		templateData.Issues = append(
			templateData.Issues,
			templates.Issue{
				Author: Author,
				Labels: Labels,
				Link:   Link,
				Number: Number,
				Title:  Title,
			},
		)
	}

	templateData.BeforeCursor = *apiData.Repository.Issues.PageInfo.GetStartCursor()
	templateData.AfterCursor = *apiData.Repository.Issues.PageInfo.GetEndCursor()
	templateData.NextPage = apiData.Repository.Issues.PageInfo.GetHasNextPage()
	templateData.PreviousPage = apiData.Repository.Issues.PageInfo.GetHasPreviousPage()
}
