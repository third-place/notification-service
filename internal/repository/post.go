package repository

import (
	"errors"
	"github.com/third-place/notification-service/internal/entity"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type PostRepository struct {
	conn *gorm.DB
}

func CreatePostRepository(conn *gorm.DB) *PostRepository {
	return &PostRepository{conn}
}

func (p *PostRepository) FindOneByUuid(uuid uuid.UUID) (*entity.Post, error) {
	post := &entity.Post{}
	p.conn.
		Preload("User").
		Where("uuid = ?", uuid).
		Find(post)
	if post.ID == 0 {
		return nil, errors.New("user not found")
	}
	return post, nil
}

func (p *PostRepository) Create(post *entity.Post) {
	p.conn.Create(post)
}

func (p *PostRepository) Save(post *entity.Post) {
	p.conn.Save(post)
}
