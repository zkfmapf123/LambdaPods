package domains

import (
	"time"
)

type Lambda struct {
	ID              uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Arn             string     `gorm:"type:varchar(255);not null;unique" json:"arn"`
	Revision        int        `gorm:"type:int;not null;default:1" json:"revision"`
	CurrentRevision int        `gorm:"type:int;not null;default:1" json:"current_revision"`
	CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

func (l Lambda) IsDelete() bool {
	return l.DeletedAt != nil
}
