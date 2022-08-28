package main

import (
	"cost-manager/lib"
)

func createConfig(mail, token string) error {
	if err := lib.ConfigInit(mail, token); err != nil {
		return err
	}

	jiraClient, _ = lib.NewJiraClient()
	jiraUser, _, _ = jiraClient.Client.User.GetSelf()

	return nil
}
