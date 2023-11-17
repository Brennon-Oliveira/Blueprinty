package commands

import (
	"Blueprinty/utils"
	"os"
)

func List(args []string) {
	files, _ := os.ReadDir(utils.GetBlueprintyDir() + "/templates")

	for _, file := range files {
		if file.IsDir() {
			println(file.Name())
		}
	}
}
