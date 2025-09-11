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
	Tracks []Track
}



type Album struct {
	Artists []string
	Title string
	Year uint64
	Discs []Disc
}





