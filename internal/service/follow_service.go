package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/db"
	"github.com/third-place/notification-service/internal/mapper"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/repository"
)

type FollowService struct {
	followRepository *repository.FollowRepository
	userRepository   *repository.UserRepository
}

func CreateFollowService() *FollowService {
	conn := db.CreateDefaultConnection()
	return &FollowService{
		repository.CreateFollowRepository(conn),
		repository.CreateUserRepository(conn),
	}
}

func (f *FollowService) UpsertFollow(followModel *model.Follow) bool {
	followUuid, err := uuid.Parse(followModel.Uuid)
	if err != nil {
		return false
	}
	userUuid, err := uuid.Parse(followModel.User.Uuid)
	if err != nil {
		return false
	}
	followEntity, err := f.followRepository.FindOneByUuid(followUuid)
	if err == nil {
		f.followRepository.Delete(followEntity)
		return false
	}
	followingUuid, err := uuid.Parse(followModel.Following.Uuid)
	if err != nil {
		return false
	}
	user, err := f.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		return false
	}
	following, err := f.userRepository.FindOneByUuid(followingUuid)
	if err != nil {
		return false
	}
	followEntity = mapper.GetFollowEntityFromModel(user.ID, following.ID, followModel)
	f.followRepository.Create(followEntity)
	return true
}
