package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)



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



//func readAlbums(folders []string) [][]Track {
func readAlbums(folders []string, exclude_folders []string) {

	//path := "./" // change this to the directory you want
	//path = "~/Music"

	frontier := expandAll(folders)
	exclude := expandAll(exclude_folders)

	for len(frontier) > 0 {
		path := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		entries, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, entry := range entries {
			fullPath := filepath.Join(path, entry.Name())
			if contains(exclude, fullPath) {
				continue;
			}

			if entry.IsDir() {
				fmt.Println("[DIR ]", fullPath)
				frontier = append(frontier, fullPath)
			} else {
				fmt.Println("[FILE]", fullPath)
			}
		}
		
	}


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
	//albums := readAlbums(req.Folders)
	albums := dummyAlbums()


    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(AlbumsResponse{Albums: albums})
}


