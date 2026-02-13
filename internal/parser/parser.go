package parser

import (
	"strings"
)

func Substitute(script string, vars map[string]string) string {
	result := script
	for k, v := range vars {
		result = strings.ReplaceAll(result, "{{"+k+"}}", v)
	}
	return result
}
