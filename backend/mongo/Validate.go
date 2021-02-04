package mongo

import (
	"encoding/hex"
)

func ValidateObjectId(id string) bool {
	d, err := hex.DecodeString(id)
	if err != nil || len(d) != 12 {
		return false
	}
	return true
}