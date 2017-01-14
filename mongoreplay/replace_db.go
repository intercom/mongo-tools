package mongoreplay

import "strings"

func replaceDB(collection string) string {
	if len(ForceDatabase) == 0 {
		return collection
	}

	if collection == "admin" || collection == "local" ||
		strings.HasPrefix(collection, "admin.") || strings.HasPrefix(collection, "local.") {
		return collection
	}

	parts := strings.Split(collection, ".")
	parts[0] = ForceDatabase
	return strings.Join(parts, ".")
}
