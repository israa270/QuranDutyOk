package common

import(
	uuid "github.com/satori/go.uuid"
)

// GetUUID get uuid
func (c *CommonUsecase) GetUUID() uuid.UUID {
	return uuid.NewV4()
}