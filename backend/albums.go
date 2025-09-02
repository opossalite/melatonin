package main



type Track struct {
	Title        string     `json:"title"`
	Artists      []string   `json:"artists"`
	Album        string     `json:"album"`
	AlbumArtists []string   `json:"album_artists"`
	Year         string     `json:"year"`
	TrackNo      string     `json:"track_no"`
	TrackCount   string     `json:"track_count"`
	CdNo         string     `json:"cd_no"`
	CdCount      string     `json:"cd_count"`
	// TODO: add cover
}



type Album struct {
	Artists []string
	Title string
	Tracks []Track
}






