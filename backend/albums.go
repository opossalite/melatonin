package main



type Track struct {
	Title        	string		`json:"title"`
	Artists      	[]string   	`json:"artists"`
	Album        	string     	`json:"album"`
	AlbumArtists 	[]string   	`json:"album_artists"`
	Year         	uint64		`json:"year"`
	Track      		uint64    	`json:"track"`
	TrackTotal   	uint64    	`json:"track_total"`
	Disc         	uint64    	`json:"disc"`
	DiscTotal      	uint64    	`json:"disc_total"`
	Duration		float64		`json:"duration"`
	Bitrate			uint64		`json:"bitrate"`
	// TODO: add cover
	Path		 string     	`json:"path"`
}



type Disc struct {
	Tracks []Track `json:"tracks"`
}



type Album struct {
	Artists []string `json:"artists"`
	Title string `json:"title"`
	Year uint64 `json:"year"`
	Discs []Disc `json:"discs"`
}


