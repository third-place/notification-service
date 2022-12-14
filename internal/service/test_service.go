package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
)

type TestService struct {
	userService         *UserService
	notificationService *NotificationService
	postService         *PostService
}

func CreateTestService() *TestService {
	return &TestService{
		userService:         CreateTestUserService(),
		notificationService: CreateTestNotificationService(),
		postService:         CreateTestPostService(),
	}
}

func (t *TestService) UpsertUser(userModel *model.User) {
	t.userService.UpsertUser(userModel)
}

func (t *TestService) CreateFollowNotification(followModel *model.Follow) {
	t.notificationService.CreateFollowNotification(followModel)
}

func (t *TestService) CreatePostLikeNotification(postLikeModel *model.PostLike) {
	t.notificationService.CreatePostLikeNotification(postLikeModel)
}

func (t *TestService) CreateReplyNotification(replyModel *model.Reply) {
	t.notificationService.CreateReplyNotification(replyModel)
}

func (t *TestService) GetNotifications(userUuid uuid.UUID, limit int) ([]*model.Notification, error) {
	return t.notificationService.GetNotifications(userUuid, limit)
}

func (t *TestService) AcknowledgeNotifications(userUuid uuid.UUID, ack *model.NotificationAcknowledgement) error {
	return t.notificationService.AcknowledgeNotifications(userUuid, ack)
}

func (t *TestService) UpsertPost(postModel *model.Post) {
	t.postService.UpsertPost(postModel)
}

func (t *TestService) UpsertReply(replyModel *model.Reply) bool {
	return t.postService.UpsertReply(replyModel)
}
