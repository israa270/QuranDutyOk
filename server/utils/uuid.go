package utils

import (
	uuid "github.com/satori/go.uuid"
)

// GetUUID get uuid
func GetUUID() uuid.UUID {
	return uuid.NewV4()
}
