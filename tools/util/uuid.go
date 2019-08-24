package util

import "github.com/google/uuid"

// GenerateUniqueID generates a unique id
func GenerateUniqueID() string {
	return uuid.Must(uuid.NewUUID()).String()
}
