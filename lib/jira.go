package lib

import (
	"github.com/andygrunwald/go-jira"
)

type JiraClient struct {
	Client *jira.Client
}

func NewJiraClient() (*JiraClient, error) {
	jiraName, err := GetConfigString("jira_name")
	if err != nil {
		return nil, err
	}

	jiraToken, err := GetConfigString("jira_token")
	if err != nil {
		return nil, err
	}

	tp := jira.BasicAuthTransport{
		Username: jiraName,
		Password: jiraToken,
	}

	jiraDomain, err := GetConfigString("jira_domain")
	if err != nil {
		return nil, err
	}

	base := "https://" + jiraDomain
	jiraClient, err := jira.NewClient(tp.Client(), base)
	if err != nil {
		panic(err)
	}

	return &JiraClient{
		Client: jiraClient,
	}, nil
}

func (jc *JiraClient) AddWorklog(issueId string, record *jira.WorklogRecord) error {
	_, _, err := jc.Client.Issue.AddWorklogRecord(issueId, record)
	if err != nil {
		return err
	}

	return nil
}
