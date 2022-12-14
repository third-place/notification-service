package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/db"
	"github.com/third-place/notification-service/internal/entity"
	"github.com/third-place/notification-service/internal/mapper"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/repository"
	"github.com/third-place/notification-service/internal/util"
	"log"
)

type NotificationService struct {
	userRepository         *repository.UserRepository
	notificationRepository *repository.NotificationRepository
	postRepository         *repository.PostRepository
}

func CreateNotificationService() *NotificationService {
	conn := db.CreateDefaultConnection()
	return &NotificationService{
		repository.CreateUserRepository(conn),
		repository.CreateNotificationRepository(conn),
		repository.CreatePostRepository(conn),
	}
}

func CreateTestNotificationService() *NotificationService {
	conn := util.SetupTestDatabase()
	return &NotificationService{
		repository.CreateUserRepository(conn),
		repository.CreateNotificationRepository(conn),
		repository.CreatePostRepository(conn),
	}
}

func (n *NotificationService) GetNotifications(userUuid uuid.UUID, limit int) ([]*model.Notification, error) {
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return nil, err
	}
	notificationEntities := n.notificationRepository.FindByUser(user, limit)
	return mapper.GetNotificationModelsFromEntities(notificationEntities), nil
}

func (n *NotificationService) AcknowledgeNotifications(userUuid uuid.UUID, ack *model.NotificationAcknowledgement) error {
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return err
	}
	result := n.notificationRepository.AcknowledgeNotifications(user.ID, ack)
	return result.Error
}

func (n *NotificationService) CreateFollowNotification(followModel *model.Follow) {
	user, err := n.userRepository.FindOneByUuid(uuid.MustParse(followModel.User.Uuid))
	if err != nil {
		return
	}
	following, err := n.userRepository.FindOneByUuid(uuid.MustParse(followModel.Following.Uuid))
	if err != nil {
		return
	}
	search, _ := n.notificationRepository.FindFollowNotification(user, following)
	if search != nil {
		return
	}
	notificationUuid := uuid.New()
	notification := &entity.Notification{
		Uuid:              &notificationUuid,
		UserID:            following.ID,
		Seen:              false,
		Link:              "/u/" + user.Username,
		NotificationType:  model.FOLLOWED,
		TriggeredByUserID: user.ID,
	}
	n.notificationRepository.Create(notification)
}

func (n *NotificationService) CreatePostLikeNotification(postLikeModel *model.PostLike) {
	userUuid, err := uuid.Parse(postLikeModel.User.Uuid)
	if err != nil {
		log.Print("error parsing userUuid :: ", err)
		return
	}
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		log.Print("user not found :: {} :: {}", userUuid, postLikeModel.Post.Uuid)
		return
	}
	postUuid, err := uuid.Parse(postLikeModel.Post.Uuid)
	if err != nil {
		return
	}
	postEntity, err := n.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		log.Print("post not found :: {}", postUuid)
		return
	}
	link := "/p/" + postLikeModel.Post.Uuid
	search, _ := n.notificationRepository.FindPostLikeNotification(user, postEntity.User, link)
	if search != nil {
		log.Print("notification already found :: ", search.Uuid)
		return
	}
	notificationUuid := uuid.New()
	notification := &entity.Notification{
		Uuid:              &notificationUuid,
		UserID:            postEntity.User.ID,
		Seen:              false,
		Link:              link,
		NotificationType:  model.POST_LIKED,
		TriggeredByUserID: user.ID,
	}
	result := n.notificationRepository.Create(notification)
	if result.Error != nil {
		log.Print("error creating notification entity :: ", result.Error)
	}
}

func (n *NotificationService) CreateReplyNotification(replyModel *model.Reply) {
	userUuid, err := uuid.Parse(replyModel.User.Uuid)
	if err != nil {
		log.Print("error parsing reply model userUuid :: ", err)
		return
	}
	user, err := n.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		log.Print("user not found :: {} :: {}", userUuid, replyModel.User.Uuid)
		return
	}
	postUuid, err := uuid.Parse(replyModel.Post.Uuid)
	if err != nil {
		return
	}
	postEntity, err := n.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		log.Print("post not found :: {}", postUuid)
		return
	}
	link := "/p/" + replyModel.Post.Uuid + "/#" + replyModel.Uuid
	search, _ := n.notificationRepository.FindReplyNotification(user, postEntity.User, link)
	if search != nil {
		log.Print("notification already found :: ", search.Uuid)
		return
	}
	notificationUuid := uuid.New()
	notification := &entity.Notification{
		Uuid:              &notificationUuid,
		UserID:            postEntity.User.ID,
		Seen:              false,
		Link:              link,
		NotificationType:  model.REPLIED,
		TriggeredByUserID: user.ID,
	}
	result := n.notificationRepository.Create(notification)
	if result.Error != nil {
		log.Print("error creating notification entity :: ", result.Error)
	}
}
