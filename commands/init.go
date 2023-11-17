package commands

import (
	"Blueprinty/utils"
	"os"
)

func Init(args []string) {

	blueprintyDir := utils.GetBlueprintyDir()

	if !utils.DirExists(blueprintyDir) {
		os.Mkdir(blueprintyDir, 0755)
	}

	if !utils.DirExists(blueprintyDir + "/templates") {
		os.Mkdir(blueprintyDir+"/templates", 0755)
	}
}
