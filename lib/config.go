package lib

import (
	"encoding/json"
	"errors"
	"github.com/rkoesters/xdg/basedir"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

const FukurouDomain = "fukurou-labo.atlassian.net/"

var (
	ConfigDir  = filepath.Join(basedir.ConfigHome, "cost-manager")
	ConfigPath = filepath.Join(ConfigDir, "config.json")
	CacheDir   = filepath.Join(basedir.CacheHome, "cost-manager")
	today      = time.Now().Format("2006-01-02")
	TodayPath  = filepath.Join(CacheDir, today+".json")
)

type Config struct {
	Mail   string `json:"mail"`
	Token  string `json:"token"`
	Domain string `json:"domain"`
}

func ConfigInit(mail, token string) error {
	if _, err := os.Stat(ConfigDir); os.IsNotExist(err) {
		if err := os.Mkdir(ConfigDir, 0777); err != nil {
			return err
		}
	}

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		fp, err := os.Create(ConfigPath)
		if err != nil {
			return err
		}
		defer fp.Close()
	}

	config := Config{
		Mail:   mail,
		Token:  token,
		Domain: FukurouDomain,
	}

	buf, err := json.Marshal(&config)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(ConfigPath, os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(buf); err != nil {
		return err
	}

	return nil
}

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
