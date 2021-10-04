package helper

import (
	"regexp"
	"strings"
)

func StringSplitter(str string) []string {
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	arr := r.FindAllString(str, -1)
	return arr
}

func RemoveQuotes(str string) string {
	return strings.Replace(str, "\"", "", -1)
}
