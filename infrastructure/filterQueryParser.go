package infrastructure

import (
	"fmt"
	"net/url"
)

func ParseFilterParams(filterQuery string) (map[string][]string, error)  {
	values, err := url.ParseQuery(filterQuery)
	if err != nil {
		return nil, fmt.Errorf("invalid filter query")
	}

	return values, nil
}
