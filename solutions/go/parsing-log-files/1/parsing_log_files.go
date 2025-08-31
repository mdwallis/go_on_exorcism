package parsinglogfiles

import (
    "fmt"
    "regexp"
)

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
    return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
    count := 0
	re := regexp.MustCompile(`(?i)".*password.*"`)
    for _, element := range lines {
        if re.MatchString(element) {
            count++
        }
    }
    return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
    updated := make([]string, len(lines))
    index := 0
    re := regexp.MustCompile(`User\s+(\w+\d+)`)
    for _, element := range lines {
        matched := re.FindStringSubmatch(element)
        if matched != nil {
            updated[index] = fmt.Sprintf("[USR] %s %s", matched[1], element)
        } else {
            updated[index] = element
        }
        index++
    }
    return updated
}
