package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
)

type TestService struct {
	userService         *UserService
	notificationService *NotificationService
}

func CreateTestService() *TestService {
	return &TestService{
		userService:         CreateTestUserService(),
		notificationService: CreateTestNotificationService(),
	}
}

func (t *TestService) UpsertUser(userModel *model.User) {
	t.userService.UpsertUser(userModel)
}

func (t *TestService) CreateFollowNotification(followModel *model.Follow) {
	t.notificationService.CreateFollowNotification(followModel)
}

func (t *TestService) GetNotifications(userUuid uuid.UUID, limit int) ([]*model.Notification, error) {
	return t.notificationService.GetNotifications(userUuid, limit)
}

func (t *TestService) AcknowledgeNotifications(userUuid uuid.UUID, ack *model.NotificationAcknowledgement) error {
	return t.notificationService.AcknowledgeNotifications(userUuid, ack)
}
