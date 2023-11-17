package commands

import (
	"Blueprinty/utils"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func Use(args []string) {

	if len(args) < 1 {
		println("Informe o template")
		return
	}

	files, _ := os.ReadDir(utils.GetBlueprintyDir() + "/templates")

	var dir string

	for _, file := range files {
		if file.IsDir() && file.Name() == args[0] {
			dir = file.Name()
			break
		}
	}

	if dir == "" {
		println("Template", args[0], "não encontrado")
		return
	}

	// exists file .config. If not exists, say that the template is not configurated
	titles, err := readConfigFile(utils.GetBlueprintyDir() + "/templates/" + dir)

	if err != nil {
		println(err.Error())
		return
	}

	userResponses := getUserValues(titles)

	createFiles(utils.GetBlueprintyDir()+"/templates/"+dir, userResponses)
}

func readConfigFile(templatePath string) ([]string, error) {
	if !utils.DirExists(templatePath + "/.config") {
		return nil, errors.New("Template não configurado")
	}

	file, err := os.Open(templatePath + "/.config")
	if err != nil {
		return nil, errors.New("Erro ao abrir arquivo de configuração do template")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	titles := make([]string, 0)
	titles = append(titles, "")

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if char == '\r' {
			continue
		}
		if char == '\n' {
			titles = append(titles, "")
			continue
		}
		titles[len(titles)-1] += string(char)
	}

	return titles, nil
}

func getUserValues(titles []string) map[string]string {

	reader := bufio.NewReader(os.Stdin)

	values := make(map[string]string)

	for _, title := range titles {
		if title == "" {
			continue
		}
		fmt.Print(title, ": ")
		value, _ := reader.ReadString('\n')
		values[title] = strings.Replace(value, "\n", "", -1)
	}

	return values
}

func createFiles(templatePath string, userResponses map[string]string) {
	files, _ := os.ReadDir(templatePath)

	for _, file := range files {
		if file.IsDir() {
			createFiles(templatePath+"/"+file.Name(), userResponses)
			continue
		}
		createFile(templatePath+"/"+file.Name(), userResponses)
	}
}

func createFile(filePath string, userResponses map[string]string) {
	file, err := os.Open(filePath)
	if err != nil {
		println("Erro ao abrir arquivo", filePath)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var content string
	var path string
	firstLine := true

	//hasPlus := false
	//hasFirstCurlyBraces := false
	//hasSecondCurlyBraces := false
	//hasFirstCloseBrackets := false

	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			break
		}
		if firstLine {
			if char == '\n' {
				firstLine = false
				continue
			}
			if char == '+' {

			}
			path += string(char)
		}

		content += string(char)
	}

	//file, err = os.Create()
	//if err != nil {
	//	println("Erro ao criar arquivo", filePath)
	//	return
	//}
	//defer file.Close()
	//
	//file.WriteString(content)
}
