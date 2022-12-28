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

func CreateTestPostService() *PostService {
	conn := util.SetupTestDatabase()
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
	postUuid, err := uuid.Parse(replyModel.Post.Uuid)
	if err != nil {
		log.Print("could not parse reply post uuid :: ", err)
		return false
	}
	postEntity, err := p.postRepository.FindOneByUuid(postUuid)
	if err != nil {
		return false
	}
	replyEntity, err := p.postRepository.FindOneByUuid(replyUuid)
	if err == nil {
		replyEntity.UpdateReplyFromModel(replyModel)
		p.postRepository.Save(replyEntity)
		return false
	}
	user, err := p.userRepository.FindOneByUuid(userUuid)
	if err != nil {
		log.Print("user not found when creating reply :: ", replyModel)
		return false
	}
	replyEntity = mapper.GetPostEntityFromReplyModel(user.ID, replyModel, postEntity.ID)
	p.postRepository.Create(replyEntity)
	return replyModel.Post.Uuid != replyModel.User.Uuid
}
