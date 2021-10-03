package helper

import (
	"regexp"
)

func StringSplitter(str string) []string {
	r := regexp.MustCompile(`[^\s"']+|"([^"]*)"|'([^']*)`)
	arr := r.FindAllString(str, -1)
	return arr
}
