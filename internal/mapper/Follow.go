package mapper

import (
	"github.com/third-place/notification-service/internal/entity"
	"github.com/third-place/notification-service/internal/model"
	"github.com/google/uuid"
)

func GetFollowEntityFromModel(userId uint, followingId uint, follow *model.Follow) *entity.Follow {
	followUuid := uuid.MustParse(follow.Uuid)
	return &entity.Follow{
		Uuid:        &followUuid,
		UserID:      userId,
		FollowingID: followingId,
	}
}
