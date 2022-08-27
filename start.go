package main

import (
	"cost-manager/lib"
	"time"
)

func start(id string) error {
	_ = lib.CacheInit()

	issue := lib.Issue{
		ID:        id,
		StartedAt: time.Now(),
	}

	if err := issue.Start(); err != nil {
		return err
	}

	return nil
}
