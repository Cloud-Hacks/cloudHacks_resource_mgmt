package filter

import (
	"fmt"

	"resource-service/utils/constants"
)

//QueryFilterString - To get the query string
func QueryFilterString(filter []string, query string) string {

	//Create the filter query
	for i, value := range filter {
		query = query + constants.BACKWARD_SLASH + fmt.Sprint(value) + constants.BACKWARD_SLASH
		if len(filter) > i+1 {
			query += constants.COMMA
		}
	}
	query += constants.CLOSE_BRACKET
	return query

}
