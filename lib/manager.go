package lib

import (
	"encoding/json"
	"errors"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"path/filepath"
	"time"
)

var (
	ConfigDir  = filepath.Join(basedir.ConfigHome, "cost-manager")
	ConfigPath = filepath.Join(ConfigDir, "config.json")
	CacheDir  = filepath.Join(basedir.CacheHome, "cost-manager")
	today     = time.Now().Format("2006-01-02")
	TodayPath = filepath.Join(CacheDir, today+".json")
)

func GetConfigString(s string) (string, error) {
	configBlob, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		return "", errors.New("command failed")
	}

	var configJson interface{}
	err = json.Unmarshal(configBlob, &configJson)
	if err != nil {
		return "", errors.New("command failed")
	}
	params := configJson.(map[string]interface{})[s].(string)

	return params, nil
}
