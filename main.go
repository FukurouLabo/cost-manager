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
	jiraUser   *jira.User
)

func init() {
	jiraClient, _ = lib.NewJiraClient()
	jiraUser, _, _ = jiraClient.Client.User.GetSelf()
}

func fetchIssueList() []jira.Issue {
	jql := "assignee=" + jiraUser.AccountID + "&status!=done"
	issues, _, _ := jiraClient.Client.Issue.SearchWithContext(context.Background(), jql, nil)
	return issues
}

func fetchRecordingIssueId() string {
	return ""
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
	app.Bind(fetchRecordingIssueId)
	app.Run()
}
