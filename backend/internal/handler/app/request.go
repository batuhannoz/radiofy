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
