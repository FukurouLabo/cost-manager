package main

import (
	"cost-manager/lib"
	_ "embed"
	"os"

	"github.com/andygrunwald/go-jira"
	"github.com/wailsapp/wails"
	"golang.org/x/net/context"
)

var (
	jiraClient *lib.JiraClient
	jiraUser   *jira.User
)

func fetchConfigFileExists() error {
	_, err := os.Stat(lib.ConfigPath)

	return err
}

func createConfigFile(mail string, token string) error {
	// TODO:Configファイルの作成・書き込み & init時に行えなかったjiraClient等の設定
	// if _, err := os.Stat(lib.ConfigDir); os.IsNotExist(err) {
	// 	if err := os.Mkdir(lib.ConfigDir, 0777); err != nil {
	// 		return err
	// 	}
	// }

	// if _, err := os.Stat(lib.ConfigPath); os.IsNotExist(err) {
	// 	fp, err := os.Create(lib.ConfigPath)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	defer fp.Close()
	// }

	return nil
}

func init() {
	err := fetchConfigFileExists()
	if err == nil {
		jiraClient, _ = lib.NewJiraClient()
		jiraUser, _, _ = jiraClient.Client.User.GetSelf()
		_ = lib.CacheInit()
	}
}

func fetchIssueList() []jira.Issue {
	jql := "assignee=" + jiraUser.AccountID + "&status!=done"
	issues, _, _ := jiraClient.Client.Issue.SearchWithContext(context.Background(), jql, nil)
	return issues
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
	app.Bind(fetchIssueList)
	app.Bind(fetchState)
	app.Bind(start)
	app.Bind(finish)
	app.Bind(fetchConfigFileExists)
	app.Bind(createConfigFile)
	app.Run()
}
