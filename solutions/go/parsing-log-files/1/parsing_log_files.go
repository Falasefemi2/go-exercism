package parsinglogfiles

import (
	"regexp"
    "strings"
)

var (
	validLineRe = regexp.MustCompile(`^\[(ERR|INF)\]`)
	splitRe     = regexp.MustCompile(`<[^a-zA-Z0-9]*>`)
)

func IsValidLine(text string) bool {
	return validLineRe.MatchString(text)
}

func SplitLogLine(text string) []string {
    	return splitRe.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	quotedRe := regexp.MustCompile(`"([^"]+)"`)
	count := 0

	for _, line := range lines {
		matches := quotedRe.FindAllStringSubmatch(line, -1)

		for _, m := range matches {
			quoted := strings.ToLower(m[1])

			if strings.Contains(quoted, "password") {
				count++
			}
		}
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	result := re.ReplaceAllString(text, "")
	return result
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	out := make([]string, 0, len(lines))

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			user := match[1]
			tagged := "[USR] " + user + " " + line
			out = append(out, tagged)
		} else {
			out = append(out, line)
		}
	}

	return out
}
