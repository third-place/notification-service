package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	Uuid        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserID      uint
	User        *User
	Following   *User
	FollowingID uint `gorm:"foreignkey:User"`
}
