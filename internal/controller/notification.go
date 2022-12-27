package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/service"
	"github.com/third-place/notification-service/internal/util"
	"net/http"
)

const notificationLimit = 100

// AcknowledgeNotificationsForUserV1 - Acknowledge notifications for a user
func AcknowledgeNotificationsForUserV1(c *gin.Context) {
	session, err := util.GetSession(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	notificationModel, _ := model.DecodeRequestToNotificationAcknowledgement(c.Request)
	err = service.CreateNotificationService().AcknowledgeNotifications(
		uuid.MustParse(session.User.Uuid),
		notificationModel,
	)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
}

// GetNotificationsForUserV1 - Get notifications for a user
func GetNotificationsForUserV1(c *gin.Context) {
	session, err := util.GetSession(c)
	if err != nil {
		c.Status(http.StatusForbidden)
		return
	}
	notifications, err := service.CreateNotificationService().
		GetNotifications(uuid.MustParse(session.User.Uuid), notificationLimit)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, notifications)
}
