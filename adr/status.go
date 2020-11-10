package adr

import (
	"fmt"
	"regexp"
)

// ChangeStatusInADRContent change the status in the ADR file content
//
func ChangeStatusInADRContent(status string, content string) (string, error) {
	re := regexp.MustCompile(`(?mi)^Status:\s?(.+)$`)
	if !re.MatchString(content) {
		return "", fmt.Errorf("file content have not a status field")
	}

	return re.ReplaceAllString(content, "Status: "+status), nil
}

// ValidateStatus validates if a status is in a list of allowed statuses
//
func ValidateStatus(targetStatus string, allowed []string) bool {
	for _, status := range allowed {
		if status == targetStatus {
			return true
		}
	}
	return false
}
