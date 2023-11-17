package utils

import "strings"

func ProcessTemplate(template string, variables map[string]string) string {
	for key, value := range variables {
		template = strings.ReplaceAll(template, "+{{"+key+"}}", value)
	}
	return template
}
