package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/db"
	"github.com/third-place/notification-service/internal/mapper"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/repository"
	"github.com/third-place/notification-service/internal/util"
	"log"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func CreateUserService() *UserService {
	return &UserService{
		repository.CreateUserRepository(db.CreateDefaultConnection()),
	}
}

func CreateTestUserService() *UserService {
	return &UserService{
		repository.CreateUserRepository(util.SetupTestDatabase()),
	}
}

func (u *UserService) UpdateProfilePic(userUuid uuid.UUID, s3Key string) {
	log.Print("update user profile pic :: {}, {}", userUuid, s3Key)
	userEntity, err := u.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		log.Print("user not found when updating profile pic")
		return
	}
	log.Print("update user with s3 key", userEntity.Uuid.String(), s3Key)
	userEntity.ProfilePic = s3Key
	u.userRepository.Save(userEntity)
}

func (u *UserService) UpsertUser(userModel *model.User) {
	userEntity, err := u.userRepository.FindOneByUuid(uuid.MustParse(userModel.Uuid))
	if err == nil {
		userEntity.UpdateUserProfileFromModel(userModel)
		u.userRepository.Save(userEntity)
	} else {
		userEntity = mapper.GetUserEntityFromModel(userModel)
		u.userRepository.Create(userEntity)
	}
}
