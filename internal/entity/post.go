package entity

import (
	"github.com/google/uuid"
	"github.com/third-place/notification-service/internal/model"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Text          string
	Draft         bool
	UserID        uint
	User          *User
	ReplyToPostID *uint `gorm:"foreignkey:Post;default:null"`
	ReplyToPost   *Post
	Visibility    model.Visibility `gorm:"default:'public'"`
	Uuid          *uuid.UUID       `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (p *Post) UpdatePostFromModel(post *model.Post) {
	p.Text = post.Text
	p.Visibility = post.Visibility
	p.Draft = post.Draft
}

func (p *Post) UpdateReplyFromModel(reply *model.Reply) {
	p.Text = reply.Text
}
