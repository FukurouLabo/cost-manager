package lib

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func CacheInit() error {
	if _, err := os.Stat(CacheDir); os.IsNotExist(err) {
		if err := os.Mkdir(CacheDir, 0777); err != nil {
			return err
		}
	}

	if _, err := os.Stat(TodayPath); os.IsNotExist(err) {
		fp, err := os.Create(TodayPath)
		if err != nil {
			return err
		}
		defer fp.Close()
	}

	return nil
}

func IssueRead(trackPath string) ([]*Issue, error) {
	jsonBytes, err := ioutil.ReadFile(trackPath)
	if err != nil {
		log.Fatal(err)
	}

	if len(jsonBytes) == 0 {
		return issues, nil
	}

	err = json.Unmarshal(jsonBytes, &issues)
	if err != nil {
		log.Fatal(err)
	}

	return issues, nil
}

func Write(trackPath string, histories []*Issue) error {
	buf, err := json.Marshal(histories)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(trackPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}
