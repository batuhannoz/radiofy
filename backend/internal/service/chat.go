package service

type ChatStore interface {
	//
}

type Chat struct {
	ChatStore
}

func NewChatService(chatStore ChatStore) *Chat {
	return &Chat{
		chatStore,
	}
}
