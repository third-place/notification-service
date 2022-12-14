package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/util"
	"testing"
	"time"
)

func Test_CanCreate_FollowNotification(t *testing.T) {
	// setup
	svc := CreateTestService()
	alice := util.CreateTestUser()
	bob := util.CreateTestUser()
	svc.UpsertUser(alice)
	svc.UpsertUser(bob)

	// given
	svc.CreateFollowNotification(&model.Follow{
		User:      *alice,
		Following: *bob,
	})

	// when
	notifications, err := svc.GetNotifications(uuid.MustParse(bob.Uuid), 1)

	// then
	if err != nil {
		t.Error(err)
	}

	if len(notifications) != 1 {
		t.Fail()
	}
}

func Test_CanAcknowledge_Notifications(t *testing.T) {
	// setup
	svc := CreateTestService()
	alice := util.CreateTestUser()
	bob := util.CreateTestUser()
	svc.UpsertUser(alice)
	svc.UpsertUser(bob)
	startAck := time.Now()

	// given
	svc.CreateFollowNotification(&model.Follow{
		User:      *alice,
		Following: *bob,
	})

	// when
	err := svc.AcknowledgeNotifications(uuid.MustParse(bob.Uuid), &model.NotificationAcknowledgement{
		DatetimeStarted: startAck,
		DatetimeEnded:   time.Now(),
	})

	// then
	if err != nil {
		t.Error(err)
	}

	// when
	notifications, err := svc.GetNotifications(uuid.MustParse(bob.Uuid), 1)

	// then
	if err != nil {
		t.Error(err)
	}

	if !notifications[0].Seen {
		t.Fail()
	}
}

func Test_CanCreate_PostLikeNotification(t *testing.T) {
	// setup
	svc := CreateTestService()
	alice := util.CreateTestUser()
	bob := util.CreateTestUser()
	svc.UpsertUser(alice)
	svc.UpsertUser(bob)
	post := &model.Post{
		Uuid: uuid.New().String(),
		User: *alice,
	}
	svc.UpsertPost(post)

	// given
	svc.CreatePostLikeNotification(&model.PostLike{
		Post: *post,
		User: *bob,
	})

	// when
	notifications, err := svc.GetNotifications(uuid.MustParse(alice.Uuid), 1)

	// then
	if err != nil {
		t.Error(err)
	}

	if len(notifications) != 1 {
		t.Fail()
	}
}

func Test_CanCreate_ReplyNotification(t *testing.T) {
	// setup
	svc := CreateTestService()
	alice := util.CreateTestUser()
	bob := util.CreateTestUser()
	svc.UpsertUser(alice)
	svc.UpsertUser(bob)
	post := &model.Post{
		Uuid: uuid.New().String(),
		User: *alice,
	}
	svc.UpsertPost(post)

	// given
	svc.CreateReplyNotification(&model.Reply{
		Uuid: uuid.New().String(),
		Post: *post,
		User: *bob,
	})

	// when
	notifications, err := svc.GetNotifications(uuid.MustParse(alice.Uuid), 1)

	// then
	if err != nil {
		t.Error(err)
	}

	if len(notifications) != 1 {
		t.Fail()
	}
}
