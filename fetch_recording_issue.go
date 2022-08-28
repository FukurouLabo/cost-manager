package main

import (
	"cost-manager/lib"
	"encoding/json"
	"io/ioutil"
	"log"
)

func fetchRecordingIssue() string {
	_ = lib.CacheInit()

	histories, err := trackRead(lib.TodayPath)
	if err != nil {
		return ""
	}

	if len(histories) == 0 {
		return ""
	}

	lastHistory := histories[len(histories)-1]
	if lastHistory.Duration > 0 {
		return ""
	}

	return lastHistory.ID
}

func trackRead(trackPath string) ([]*lib.Issue, error) {
	jsonBytes, err := ioutil.ReadFile(trackPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return lib.Issues, nil
	}

	err = json.Unmarshal(jsonBytes, &lib.Issues)
	if err != nil {
		log.Fatal(err)
	}

	return lib.Issues, nil
}
