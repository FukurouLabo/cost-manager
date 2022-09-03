package main

import (
	"cost-manager/lib"
	"errors"
)

func cancelMeasurement() error {
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

	newHistories := histories[:len(histories)-1]

	if err := lib.Write(lib.TodayPath, newHistories); err != nil {
		return err
	}

	return nil
}
