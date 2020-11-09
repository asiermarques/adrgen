package adr

import (
	"fmt"
	"regexp"
)

func ChangeStatusInADRContent(status string, content string) (string, error) {
	re := regexp.MustCompile(`(?mi)^Status:\s?(.+)$`)
	if !re.MatchString(content) {
		return "", fmt.Errorf("file content have not a status field")
	}

	return re.ReplaceAllString(content, "Status: "+status), nil
}

func ValidateStatus(targetStatus string, allowed []string) bool {
	for _, status := range allowed {
		if status == targetStatus {
			return true
		}
	}
	return false
}
