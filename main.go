package main

import (
	"Blueprinty/commands"
	"Blueprinty/utils"
	"fmt"
	"os"
)

func Help() {
	fmt.Println("Help")
}

func main() {

	funcMap := make(map[string]func(args []string))
	funcMap["init"] = commands.Init
	funcMap["list"] = commands.List
	funcMap["use"] = commands.Use

	if len(os.Args) < 2 {
		Help()
		return
	}
	arg := os.Args[1]
	args := os.Args[2:]

	projectIsConfigurated := utils.ProjectIsConfigurated()

	if arg != "init" && !projectIsConfigurated {
		fmt.Println("Projeto ainda não configurado. Execute 'blueprinty init' para inicializar o Blueprinty")
		return
	}

	if f, ok := funcMap[arg]; ok {
		f(args)
	} else {
		fmt.Println("Comando", arg, "não encontrado")
	}
}
