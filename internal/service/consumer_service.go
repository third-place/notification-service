package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
)

type ConsumerService struct {
	followService       *FollowService
	notificationService *NotificationService
	postService         *PostService
	userService         *UserService
}

func CreateConsumerService() *ConsumerService {
	return &ConsumerService{
		CreateFollowService(),
		CreateNotificationService(),
		CreatePostService(),
		CreateUserService(),
	}
}

func (c *ConsumerService) UpsertUser(userModel *model.User) {
	c.userService.UpsertUser(userModel)
}

func (c *ConsumerService) UpdateProfilePic(userUuid uuid.UUID, s3Key string) {
	c.userService.UpdateProfilePic(userUuid, s3Key)
}

func (c *ConsumerService) UpsertPost(postModel *model.Post) {
	c.postService.UpsertPost(postModel)
}

func (c *ConsumerService) UpsertReply(replyModel *model.Reply) {
	shouldSendNotification := c.postService.UpsertReply(replyModel)
	if shouldSendNotification {
		c.notificationService.CreateReplyNotification(replyModel)
	}
}

func (c *ConsumerService) UpsertFollow(followModel *model.Follow) {
	shouldSendNotification := c.followService.UpsertFollow(followModel)
	if shouldSendNotification {
		c.notificationService.CreateFollowNotification(followModel)
	}
}

func (c *ConsumerService) CreatePostLikeNotification(postLike *model.PostLike) {
	c.notificationService.CreatePostLikeNotification(postLike)
}
