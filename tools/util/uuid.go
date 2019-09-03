package util

import "github.com/google/uuid"

func GenerateUniqueID() string {
	return uuid.Must(uuid.NewUUID()).String()
}
