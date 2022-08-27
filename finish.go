package main

import (
	"cost-manager/lib"
	"errors"
	"time"
)

func finish() error {
	histories, err := lib.IssueRead(lib.TodayPath)
	if err != nil {
		return err
	}

	if len(histories) == 0 {
		return errors.New("task is not running")
	}

	lastHistory := histories[len(histories)-1]
	if lastHistory.Duration > 0 {
		return errors.New("task is not running")
	}

	now := time.Now()
	lastHistory.FinishedAt = now
	lastHistory.Duration = now.Sub(lastHistory.StartedAt)

	if err := lastHistory.Finish(jiraClient); err != nil {
		return err
	}

	return nil
}
