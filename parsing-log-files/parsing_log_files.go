package parsinglogfiles

import (
	"fmt"
	"regexp"
)

var validationLine = `^\[ERR\]|^\[TRC\]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\]`

func IsValidLine(text string) bool {
	validation := regexp.MustCompile(validationLine)

	return validation.MatchString(text)
}

var validationLogLine = `<[~*=-]*>`

func SplitLogLine(text string) []string {
	validation := regexp.MustCompile(validationLogLine)

	return validation.Split(text, -1)
}

var validationQuotedPasswords = `(?i)"([^"\\]*(\\.[^"\\]*)*)password([^"\\]*(\\.[^"\\]*)*)"`

func CountQuotedPasswords(lines []string) int {
	validation := regexp.MustCompile(validationQuotedPasswords)
	counter := 0

	for _, line := range lines {
		isValidLine := validation.MatchString(line)

		if isValidLine {
			counter++
		}
	}

	return counter
}

var validationEndOfLine = `end-of-line+\S*`

func RemoveEndOfLineText(text string) string {
	validation := regexp.MustCompile(validationEndOfLine)

	return validation.ReplaceAllString(text, "")
}

var validationTagWithUserName = `User+\s+[[:alnum:]]*`
var validationWhitespaces = `\s+`
var userTag = "[USR]"

func TagWithUserName(lines []string) []string {
	validation := regexp.MustCompile(validationTagWithUserName)
	validationSpaces := regexp.MustCompile(validationWhitespaces)

	for index, line := range lines {
		isMatch := validation.MatchString(line)

		if isMatch {
			matchString := validation.FindString(line)
			username := validationSpaces.Split(matchString, -1)[1]

			lines[index] = fmt.Sprintf("%s %s %s", userTag, username, line)
		}
	}

	return lines
}
