package service

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/db"
	"github.com/third-place/notification-service/internal/mapper"
	"github.com/third-place/notification-service/internal/model"
	"github.com/third-place/notification-service/internal/repository"
	"log"
)

type PostService struct {
	userRepository *repository.UserRepository
	postRepository *repository.PostRepository
}

func CreatePostService() *PostService {
	conn := db.CreateDefaultConnection()
	return &PostService{
		repository.CreateUserRepository(conn),
		repository.CreatePostRepository(conn),
	}
}

func (p *PostService) UpsertPost(postModel *model.Post) {
	postUuid, err := uuid.Parse(postModel.Uuid)
	if err != nil {
		return
	}
	userUuid, err := uuid.Parse(postModel.User.Uuid)
	if err != nil {
		return
	}
	postEntity, err := p.postRepository.FindOneByUuid(postUuid)
	if err == nil {
		postEntity.UpdatePostFromModel(postModel)
		p.postRepository.Save(postEntity)
	} else {
		user, err := p.userRepository.FindOneByUuid(userUuid)
		if err != nil {
			log.Print("user not found when upserting post :: ", postModel)
			return
		}
		postEntity = mapper.GetPostEntityFromPostModel(user.ID, postModel)
		p.postRepository.Create(postEntity)
	}
}

func (p *PostService) UpsertReply(replyModel *model.Reply) bool {
	replyUuid, err := uuid.Parse(replyModel.Uuid)
	if err != nil {
		return false
	}
	userUuid, err := uuid.Parse(replyModel.User.Uuid)
	if err != nil {
		return false
	}
	postEntity, err := p.postRepository.FindOneByUuid(replyUuid)
	if err == nil {
		postEntity.UpdateReplyFromModel(replyModel)
		p.postRepository.Save(postEntity)
		return false
	}
	user, err := p.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		log.Print("user not found when creating reply :: ", replyModel)
		return false
	}
	postEntity = mapper.GetPostEntityFromReplyModel(user.ID, replyModel, postEntity.ID)
	p.postRepository.Create(postEntity)
	return replyModel.Post.Uuid != replyModel.User.Uuid
}
