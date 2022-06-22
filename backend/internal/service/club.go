package service

type ClubStore interface {
	//
}

type Club struct {
	ChatStore
}

func NewClubService(clubStore ClubStore) *Club {
	return &Club{
		clubStore,
	}
}
