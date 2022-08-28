package main

import (
	"github.com/andygrunwald/go-jira"
	"golang.org/x/net/context"
)

func fetchIssues() []jira.Issue {
	jql := "assignee=" + jiraUser.AccountID + "&status!=done"
	issues, _, _ := jiraClient.Client.Issue.SearchWithContext(context.Background(), jql, nil)
	return issues
}
