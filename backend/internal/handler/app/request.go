package app

type CreateClubRequest struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

type PlaybackRequest struct {
	AlbumID    string `json:"albumID"`
	Position   int    `json:"position"`
	SongID     string `json:"songID"`
	SongName   string `json:"songName"`
	ArtistName string `json:"artistName"`
	Image      string `json:"image"`
	ProgressMS int    `json:"progressMS"`
}

type MessageRequest struct {
	Message string `json:"message"`
}

type CurrentListenersRequest struct {
	Position   int    `json:"position"`
	AlbumID    string `json:"albumID"`
	ArtistName string `json:"artistName"`
	SongName   string `json:"songName"`
	SongID     string `json:"songID"`
}
