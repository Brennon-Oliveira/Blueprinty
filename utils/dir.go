package utils

import (
	"os"
	"os/user"
)

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

func GetBlueprintyDir() string {
	return GetHomeDir() + "/.blueprinty"
}

func DirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func ProjectIsConfigurated() bool {
	blueprintyDir := GetBlueprintyDir()
	if !DirExists(blueprintyDir) {
		return false
	}
	if !DirExists(blueprintyDir + "/templates") {
		return false
	}
	return true
}
