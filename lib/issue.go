package lib

import (
	"time"
	"github.com/andygrunwald/go-jira"
)

var (
	Issues []*Issue
)

type Issue struct {
	ID         string           `json:"id"`
	Name       string        `json:"name"`
	StartedAt  time.Time     `json:"started_at"`
	FinishedAt time.Time     `json:"finished_at"`
	Duration   time.Duration `json:"duration"`
}

func (issue *Issue)Start() error {
	histories, err := IssueRead(TodayPath)
	if err != nil {
		return err
	}
	newHistories := append(histories, issue)

	if err := Write(TodayPath, newHistories); err != nil {
		return err
	}

	return nil
}

func (issue Issue)Finish(client *JiraClient) error {
	histories, err := IssueRead(TodayPath)
	if err != nil {
		return err
	}

	newHistories := append(histories[:len(histories)-1], &issue)

	if err := Write(TodayPath, newHistories); err != nil {
		return err
	}

		now := time.Now()
		started := issue.StartedAt
		record := jira.WorklogRecord{
			Created:          (*jira.Time)(&now),
			Updated:          (*jira.Time)(&now),
			Started:          (*jira.Time)(&started),
			TimeSpentSeconds: int(issue.Duration.Seconds()),
		}

		if err := client.AddWorklog(issue.ID, &record); err != nil {
			return err
		}

	return nil
}
