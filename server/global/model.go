package global

import (
	"time"

	"gorm.io/gorm"
)

type GvaModel struct {
	ID        uint `gorm:"primarykey" json:",omitempty"`// primary key ID
	CreatedAt time.Time   `json:",omitempty"`
	UpdatedAt time.Time    `json:",omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
