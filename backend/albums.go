package main



type Track struct {
	Title        string     `json:"title"`
	Artists      []string   `json:"artists"`
	Album        string     `json:"album"`
	AlbumArtists []string   `json:"album_artists"`
	Year         uint     `json:"year"`
	TrackNo      uint     `json:"track_no"`
	TrackCount   uint     `json:"track_count"`
	CdNo         uint     `json:"cd_no"`
	CdCount      uint     `json:"cd_count"`
	// TODO: add cover
}



type Album struct {
	Artists []string
	Title string
	Tracks []Track
}






