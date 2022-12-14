package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/third-place/notification-service/internal/entity"
	"github.com/third-place/notification-service/internal/model"
)

type NotificationRepository struct {
	conn *gorm.DB
}

func CreateNotificationRepository(conn *gorm.DB) *NotificationRepository {
	return &NotificationRepository{conn}
}

func (n *NotificationRepository) FindPostLikeNotification(user *entity.User, postUser *entity.User, link string) (*entity.Notification, error) {
	notification := &entity.Notification{}
	n.conn.
		Table("notifications").
		Where(
			"notifications.user_id = ? AND notifications.triggered_by_user_id = ? AND notifications.link = ?",
			postUser.ID,
			user.ID,
			link,
		).Find(notification)
	if notification.ID == 0 {
		return nil, errors.New("notification not found")
	}
	return notification, nil
}

func (n *NotificationRepository) FindReplyNotification(user *entity.User, postUser *entity.User, link string) (*entity.Notification, error) {
	notification := &entity.Notification{}
	n.conn.
		Table("notifications").
		Where(
			"notifications.user_id = ? AND notifications.triggered_by_user_id = ? AND notifications.link = ?",
			postUser.ID,
			user.ID,
			link,
		).Find(notification)
	if notification.ID == 0 {
		return nil, errors.New("notification not found")
	}
	return notification, nil
}

func (n *NotificationRepository) FindFollowNotification(user *entity.User, following *entity.User) (*entity.Notification, error) {
	notification := &entity.Notification{}
	n.conn.
		Table("notifications").
		Where(
			"notifications.user_id = ? AND notifications.triggered_by_user_id = ? AND notifications.notification_type = ?",
			following.ID,
			user.ID,
			model.FOLLOWED,
		).Find(notification)
	if notification.ID == 0 {
		return nil, errors.New("notification not found")
	}
	return notification, nil
}

func (n *NotificationRepository) FindByUser(user *entity.User, limit int) []*entity.Notification {
	var notifications []*entity.Notification
	n.conn.
		Preload("User").
		Preload("TriggeredByUser").
		Table("notifications").
		Where("notifications.user_id = ?", user.ID).
		Limit(limit).
		Order("id desc").
		Find(&notifications)
	return notifications
}

func (n *NotificationRepository) AcknowledgeNotifications(userID uint, ack *model.NotificationAcknowledgement) *gorm.DB {
	return n.conn.
		Model(&entity.Notification{}).
		Where("user_id = ? AND created_at >= ? AND created_at <= ?", userID, ack.DatetimeStarted, ack.DatetimeEnded).
		Update("seen", true)
}

func (n *NotificationRepository) Create(notification *entity.Notification) *gorm.DB {
	return n.conn.Create(notification)
}
