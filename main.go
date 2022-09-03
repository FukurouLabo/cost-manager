package main

import (
	"cost-manager/lib"
	_ "embed"

	"github.com/andygrunwald/go-jira"
	"github.com/wailsapp/wails"
)

var (
	jiraClient *lib.JiraClient
	jiraUser   *jira.User
)

func init() {
	if err := fetchConfig(); err == nil {
		jiraClient, _ = lib.NewJiraClient()
		jiraUser, _, _ = jiraClient.Client.User.GetSelf()
	}
	_ = lib.CacheInit()
}

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {
	app := wails.CreateApp(&wails.AppConfig{
		Width:  800,
		Height: 1024,
		Title:  "cost-manager",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})

	app.Bind(fetchIssues)
	app.Bind(fetchRecordingIssue)
	app.Bind(startMeasurement)
	app.Bind(finishMeasurement)
	app.Bind(cancelMeasurement)
	app.Bind(fetchConfig)
	app.Bind(createConfig)
	app.Run()
}
