package main

import (
	"cost-manager/lib"
	"errors"
	"time"
)

const (
	LayoutStandard = "2006-01-02T15:04:05"
	LayoutMicro    = "2006-01-02T15:04:05.999999+09:00"
)

var jst, _ = time.LoadLocation("Asia/Tokyo")

func interruptMeasurement(finishedAtStr string) error {
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

	finishedAtTime, err := time.ParseInLocation(LayoutStandard, finishedAtStr, jst)
	if err != nil {
		return err
	}
	finishedAtTime.Format(LayoutMicro)

	lastHistory.FinishedAt = finishedAtTime
	lastHistory.Duration = finishedAtTime.Sub(lastHistory.StartedAt)

	if err := lastHistory.Finish(jiraClient); err != nil {
		return err
	}

	return nil
}
