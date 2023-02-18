package _infrastructure

import (
	"github.com/asiermarques/adrgen/internal/adr"
)

const statusFilterParam = "status"

func FilterADR(adr adr.ADR, filterParams map[string][]string) bool {
	if len(filterParams[statusFilterParam]) > 0 {
		for _, status := range filterParams[statusFilterParam] {
			if status == adr.Status() {
				return true
			}
		}
	}
	return false
}
