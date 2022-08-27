package main

import (
	"cost-manager/lib"
	"time"
)

func start() error {
	_ = lib.CacheInit()
	issues := fetchIssueList()
	i := issues[0]
	issue := lib.Issue{
		ID:        i.ID,
		Name:      i.Fields.Summary,
		StartedAt: time.Now(),
	}

	if err := issue.Start(); err != nil {
		return err
	}

	return nil
}
