package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/service"
	"net/http"
)

const notificationLimit = 100

// AcknowledgeNotificationsForUserV1 - Acknowledge notifications for a user
func AcknowledgeNotificationsForUserV1(w http.ResponseWriter, r *http.Request) {
	session := service.CreateDefaultAuthService().GetSessionFromRequest(r)
	if session == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	notificationModel, _ := model.DecodeRequestToNotificationAcknowledgement(r)
	err := service.CreateNotificationService().AcknowledgeNotifications(
		uuid.MustParse(session.User.Uuid),
		notificationModel,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// GetNotificationsForUserV1 - Get notifications for a user
func GetNotificationsForUserV1(w http.ResponseWriter, r *http.Request) {
	session := service.CreateDefaultAuthService().GetSessionFromRequest(r)
	if session == nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	notifications, err := service.CreateNotificationService().
		GetNotifications(uuid.MustParse(session.User.Uuid), notificationLimit)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data, err := json.Marshal(notifications)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}
