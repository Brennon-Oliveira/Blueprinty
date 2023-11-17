package utils

import (
	"os"
	"os/user"
	"path/filepath"
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

func MakeDirs(path string) error {

	// Separa o caminho em diretório e arquivo
	dir, _ := filepath.Split(path)

	// Verifica se diretório pai existe
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// Cria o diretório pai recursivamente
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
