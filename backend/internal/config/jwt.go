package config

import "time"

type JWT struct {
	Secret string        `json:"secret"`
	Expire time.Duration `json:"expire"`
}
