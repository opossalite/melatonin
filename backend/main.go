package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Track struct {
	Title        string     `json:"title"`
	Artists      []string `json:"artists"`
	Album        string   `json:"album"`
	AlbumArtists []string `json:"album_artists"`
	Year         string     `json:"year"`
	TrackNo      string     `json:"track_no"`
	TrackCount   string     `json:"track_count"`
	CdNo         string     `json:"cd_no"`
	CdCount      string     `json:"cd_count"`
	// TODO: add cover
}

type AlbumsResponse struct {
	Albums [][]Track `json:"albums"`
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	// Example data — replace with your real lookup.
	albums := [][]Track{
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
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(AlbumsResponse{Albums: albums})
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// CORS for localhost dev (adjust as needed)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "http://localhost:8080"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/get_albums", getAlbums)

	log.Println("listening on :8800")
	log.Fatal(http.ListenAndServe(":8800", r))
}

