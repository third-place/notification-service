package model

import "time"

type Notification struct {
	Uuid             string           `json:"uuid"`
	CreatedAt        time.Time        `json:"created_at,omitempty"`
	User             User             `json:"user"`
	Seen             bool             `json:"seen"`
	Link             string           `json:"link"`
	NotificationType NotificationType `json:"notificationType,omitempty"`
	TriggeredByUser  User             `json:"triggered_by_user,omitempty"`
}
