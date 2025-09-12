package main



import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)



var PERMITTED_FILE_TYPES []string = []string{
	"mp3",

	"acc",
	"m4a",
	"alac",

	"ogg",
	"flac",
}



func expandPath(path string) (string, error) {
    if len(path) > 0 && path[0] == '~' {
        home, err := os.UserHomeDir()
        if err != nil {
            return "", err
        }
        if path == "~" {
            return home, nil
        }
        // replace "~" with the home directory
        return filepath.Join(home, path[1:]), nil
    }
    return path, nil
}



func expandAll(paths []string) []string {
    expanded := make([]string, len(paths))
    for i, p := range paths {
        e, err := expandPath(p)
        if err != nil {
            fmt.Println("Error expanding path:", p, err)
            os.Exit(1) // terminate program
        }
        expanded[i] = e
    }
    return expanded
}



func main() {
	// note, this returns a list of tracks
	//a := sortTracks(readTracks([]string{"~/Music"}, []string{"~/Music/,OLD"}))
	//fmt.Println(a)
	//return

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// CORS for localhost dev (adjust as needed)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "http://localhost:8080"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Origin"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Post("/get_albums", getAlbums)

	log.Println("listening on :8800")
	log.Fatal(http.ListenAndServe(":8800", router))
}


