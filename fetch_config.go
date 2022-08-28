package main

import (
	"cost-manager/lib"
	"os"
)

func fetchConfig() error {
	_, err := os.Stat(lib.ConfigPath)

	return err
}
