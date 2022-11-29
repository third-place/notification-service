package mapper

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/entity"
	"github.com/third-place/notification-service/internal/model"
)

func GetPostEntityFromPostModel(userId uint, post *model.Post) *entity.Post {
	postUuid := uuid.MustParse(post.Uuid)
	return &entity.Post{
		Uuid:       &postUuid,
		Text:       post.Text,
		Visibility: post.Visibility,
		Draft:      post.Draft,
		UserID:     userId,
	}
}

func GetPostEntityFromReplyModel(userId uint, reply *model.Reply, postId uint) *entity.Post {
	replyUuid := uuid.MustParse(reply.Uuid)
	return &entity.Post{
		Uuid:          &replyUuid,
		Text:          reply.Text,
		UserID:        userId,
		ReplyToPostID: postId,
	}
}
