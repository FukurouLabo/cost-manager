package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"cost-manager/lib"
)

func getState() (string, error) {
	histories, err := trackRead(lib.TodayPath)
	if err != nil {
		return "", err
	}

	if len(histories) == 0 {
		return "", nil
	}

	lastHistory := histories[len(histories)-1]
	if lastHistory.Duration > 0 {
		return "", nil
	}

	return lastHistory.ID, nil
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