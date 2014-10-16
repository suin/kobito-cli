package main

import (
	"errors"
	"os"
	"os/user"
	"strings"
)

func kobitoPath() (path string, err error) {
	usr, _ := user.Current()
	home := usr.HomeDir

	lookUpPaths := []string{
		home + "/Applications/Kobito.app",
		"/Applications/Kobito.app",
	}

	for _, path = range lookUpPaths {
		if directoryExists(path) {
			return
		}
	}

	err = errors.New("Kobito.app not found in: " + strings.Join(lookUpPaths, ", "))

	return
}

func directoryExists(dir string) bool {
	info, err := os.Stat(dir)
	if err != nil {
		return false
	}
	return info.IsDir()
}
