package helper

import (
	"regexp"
	"strconv"
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

func IsEligibleConvertToInteger(str string) bool {
	_, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return true
}
