package main

import (
	"cost-manager/lib"
	_ "embed"
	"github.com/andygrunwald/go-jira"
	"github.com/wailsapp/wails"
	"golang.org/x/net/context"
)

var (
	jiraClient *lib.JiraClient
)

func init() {
	jiraClient, _ = lib.NewJiraClient()
}

func fetchIssueList() []jira.Issue {
	// TODO: jql は変更する必要あり
	issues, _, _ := jiraClient.Client.Issue.SearchWithContext(context.Background(), "assignee!=aaa", nil)
	return issues
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "cost-manager",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(fetchIssueList)
	app.Run()

}
