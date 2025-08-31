package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)



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



func dummyAlbums() [][]Track {
	return [][]Track{
		{
			{
				Title:        "Song A",
				Artists:      []string{"Artist One", "Feat Person"},
				Album:        "Greatest Hits",
				AlbumArtists: []string{"Artist One"},
				Year:         "2024",
				TrackNo:      "1",
				TrackCount:   "10",
				CdNo:         "1",
				CdCount:      "1",
			},
			{
				Title:        "Song B",
				Artists:      []string{"Artist One"},
				Album:        "Greatest Hits",
				AlbumArtists: []string{"Artist One"},
				Year:         "2024",
				TrackNo:      "2",
				TrackCount:   "10",
				CdNo:         "1",
				CdCount:      "1",
			},
		},
		{
			{
				Title:        "Another Song",
				Artists:      []string{"Different Band"},
				Album:        "Live at Home",
				AlbumArtists: []string{"Different Band"},
				Year:         "2019",
				TrackNo:      "1",
				TrackCount:   "8",
				CdNo:         "1",
				CdCount:      "1",
			},
		},
	};

}



func readAlbums(folders []string) [][]Track {
	fmt.Println("hi")

	fmt.Println("returning dummy values")
	return dummyAlbums()
}



type AlbumsRequest struct {
    Folders []string `json:"folders"`
}
type AlbumsResponse struct {
	Albums [][]Track `json:"albums"`
}
func getAlbums(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    var req AlbumsRequest
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()
    if err := dec.Decode(&req); err != nil {
        http.Error(w, "invalid JSON: "+err.Error(), http.StatusBadRequest)
        return
    }
	albums := readAlbums(req.Folders)


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(AlbumsResponse{Albums: albums})
}


