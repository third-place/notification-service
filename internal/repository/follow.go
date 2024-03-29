package repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/entity"
	"gorm.io/gorm"
)

type FollowRepository struct {
	conn *gorm.DB
}

func CreateFollowRepository(conn *gorm.DB) *FollowRepository {
	return &FollowRepository{conn}
}

func (f *FollowRepository) Create(follow *entity.Follow) {
	f.conn.Create(follow)
}

func (f *FollowRepository) Save(follow *entity.Follow) {
	f.conn.Save(follow)
}

func (f *FollowRepository) FindOneByUuid(uuid uuid.UUID) (*entity.Follow, error) {
	follow := &entity.Follow{}
	f.conn.Where("uuid = ?", uuid).Find(follow)
	if follow.ID == 0 {
		return nil, errors.New("user not found")
	}
	return follow, nil
}

func (f *FollowRepository) Delete(follow *entity.Follow) {
	f.conn.Delete(follow)
}
