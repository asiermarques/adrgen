package infrastructure

import "github.com/asiermarques/adrgen/domain"

const statusFilterParam = "status"

func FilterADR(adr domain.ADR, filterParams map[string][]string) bool {
	if len(filterParams[statusFilterParam]) > 0 {
		for _, status := range filterParams[statusFilterParam] {
			if status == adr.Status() {
				return true
			}
		}
	}
	return false
}
