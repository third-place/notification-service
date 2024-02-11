package entity

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	Uuid              *uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4()"`
	Seen              bool                   `gorm:"default:false"`
	Link              string                 `gorm:"not null"`
	NotificationType  model.NotificationType `gorm:"index"`
	UserID            uint                   `gorm:"index"`
	User              *User
	TriggeredByUserID uint `gorm:"index"`
	TriggeredByUser   *User
}
