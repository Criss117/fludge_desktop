package dbutils

import "strings"

func IsUniqueConstraint(err error, index string) bool {
	return strings.Contains(err.Error(), index)
}
