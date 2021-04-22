package uuid

import (
	guuid "github.com/google/uuid"
)

func GenerateUUID() string {
	return guuid.NewString()
}
