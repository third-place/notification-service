package model

import (
	"encoding/json"
	"time"
)

type Reply struct {
	Uuid      string    `json:"uuid"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Post      Post      `json:"post"`
	Text      string    `json:"text"`
	User      User      `json:"user"`
}

func DecodeMessageToReply(message []byte) (*Reply, error) {
	post := &Reply{}
	err := json.Unmarshal(message, post)
	return post, err
}
