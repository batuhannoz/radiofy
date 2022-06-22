package service

type LogStore interface {
	//
}

type Log struct {
	LogStore
}

func NewLogService(logStore LogStore) *Log {
	return &Log{
		logStore,
	}
}
