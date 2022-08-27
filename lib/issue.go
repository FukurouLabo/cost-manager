package lib

import (
	"time"
)

var (
	issues []*Issue
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
