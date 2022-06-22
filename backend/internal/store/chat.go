package store

import "gorm.io/gorm"

type ChatStore struct {
	connection *gorm.DB
}

func NewChatStore(connection *gorm.DB) *ChatStore {
	return &ChatStore{
		connection,
	}
}
