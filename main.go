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
	_ = lib.CacheInit()
}

func fetchIssueList() []jira.Issue {
	jql := "assignee=" + jiraUser.AccountID + "&status!=done"
	issues, _, _ := jiraClient.Client.Issue.SearchWithContext(context.Background(), jql, nil)
	return issues
}

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "cost-manager",
	})
	app.Bind(fetchIssueList)
	app.Bind(fetchState)
	app.Bind(start)
	app.Bind(finish)
	app.Run()
}
