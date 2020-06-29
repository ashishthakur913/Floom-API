package postgres

import "fmt"

func mergeQuery(query string, condition string) string {
	if len(query) == 0 {
		return condition
	}

	return fmt.Sprintf("%s and %s", query, condition)
}
