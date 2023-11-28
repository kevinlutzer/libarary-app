package shared

import (
	"regexp"
)

var rx = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

func IsValidID(id string) bool {
	return rx.MatchString(id)
}

const InvalidIdMsg = "id specified is not valid, it must be a valid v4 uuid like 0b4c6db1-ea64-46f7-8a0f-d82f5de905a9"
