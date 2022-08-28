package main

import (
	"cost-manager/lib"

	"github.com/andygrunwald/go-jira"
	"github.com/wailsapp/wails"
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
	_ = app.Run()
}
