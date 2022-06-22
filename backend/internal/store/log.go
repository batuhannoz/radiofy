package store

import "gorm.io/gorm"

type LogStore struct {
	connection *gorm.DB
}

func NewLogStore(connection *gorm.DB) *LogStore {
	return &LogStore{
		connection,
	}
}
