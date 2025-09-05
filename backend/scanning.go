package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"time"
)



func dummyAlbums() [][]Track {
	return [][]Track{
		{
			{
				Title:        "Song A",
				Artists:      []string{"Artist One", "Feat Person"},
				Album:        "Greatest Hits",
				AlbumArtists: []string{"Artist One"},
				Year:         2024,
				TrackNo:      1,
				TrackCount:   10,
				CdNo:         1,
				CdCount:      1,
			},
			{
				Title:        "Song B",
				Artists:      []string{"Artist One"},
				Album:        "Greatest Hits",
				AlbumArtists: []string{"Artist One"},
				Year:         2024,
				TrackNo:      2,
				TrackCount:   10,
				CdNo:         1,
				CdCount:      1,
			},
		},
		{
			{
				Title:        "Another Song",
				Artists:      []string{"Different Band"},
				Album:        "Live at Home",
				AlbumArtists: []string{"Different Band"},
				Year:         2019,
				TrackNo:      1,
				TrackCount:   8,
				CdNo:         1,
				CdCount:      1,
			},
		},
	};

}



// Run ffprobe in the command line for a given file path and return a Track
func FFProbeTags(ctx context.Context, path string) (Track, error) {
    if _, ok := ctx.Deadline(); !ok {
        var cancel context.CancelFunc
        ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
        defer cancel()
    }

    // ask ffprobe to show its normal human-readable block (to stderr)
    cmd := exec.CommandContext(ctx, "ffprobe", "-hide_banner", path)

	// set up an stdout and an stderr
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    if err := cmd.Run(); err != nil {
        // If ffprobe printed something, surface it in the error
        msg := strings.TrimSpace(stderr.String())
        if msg == "" {
            msg = err.Error()
        }
        return Track{}, errors.New("ffprobe: " + msg)
    }

    sc := bufio.NewScanner(bytes.NewReader(stderr.Bytes()))
	var track Track
    for sc.Scan() {

		// set up a pipeline to extract valuable info
        line := strings.TrimSpace(sc.Text())
		lineLower := strings.ToLower(line)
        if line == "" {
			continue
        }

		if strings.HasPrefix(lineLower, "title") {
			loc := strings.Index(line, ":")
			track.Title = line[loc+2:]
			continue
		}

		if strings.HasPrefix(lineLower, "artist") {
			loc := strings.Index(line, ":")
			substr := line[loc+2:] //contains multiple artists possibly
			track.Artists = strings.Split(substr, ";")
			continue
		}

		if strings.HasPrefix(lineLower, "album_artist") {
			loc := strings.Index(line, ":")
			substr := line[loc+2:] //contains multiple artists possibly
			track.AlbumArtists = strings.Split(substr, ";")
			continue
		}

		if strings.HasPrefix(lineLower, "album") {
			loc := strings.Index(line, ":")
			track.Album = line[loc+2:]
			continue
		}

		if strings.HasPrefix(lineLower, "disctotal") {
			loc := strings.Index(line, ":")
			num, err := strconv.ParseUint(line[loc+2:], 10, 32)
			if err != nil {
				fmt.Println(err)
				continue
			}
			track.CdCount = uint(num)
			continue
		}

		if strings.HasPrefix(lineLower, "disc") {
			loc := strings.Index(line, ":")
			num, err := strconv.ParseUint(line[loc+2:], 10, 32)
			if err != nil {
				fmt.Println(err)
				continue
			}
			track.CdNo = uint(num)
			continue
		}


		if strings.HasPrefix(lineLower, "date") {
			loc := strings.Index(line, ":")
			num, err := strconv.ParseUint(line[loc+2:], 10, 32)
			if err != nil {
				fmt.Println(err)
				continue
			}
			track.Year = uint(num)
			continue
		}

		if strings.HasPrefix(lineLower, "tracktotal") {
			loc := strings.Index(line, ":")
			num, err := strconv.ParseUint(line[loc+2:], 10, 32)
			if err != nil {
				fmt.Println(err)
				continue
			}
			track.TrackCount = uint(num)
			continue
		}

		if strings.HasPrefix(lineLower, "track") {
			loc := strings.Index(line, ":")
			num, err := strconv.ParseUint(line[loc+2:], 10, 32)
			if err != nil {
				fmt.Println(err)
				continue
			}
			track.TrackNo = uint(num)
			continue
		}
    }
    if err := sc.Err(); err != nil {
        return Track{}, err
    }

    return track, nil
}



//func readAlbums(folders []string) [][]Track {
func readAlbums(folders []string, exclude_folders []string) {
	frontier := expandAll(folders)
	exclude := expandAll(exclude_folders)

	tracks := []Track{}; //will be sorted later
	failed := []string{}

	ctx := context.Background() //for commandline use

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
			if slices.Contains(exclude, fullPath) {
				continue;
			}

			if entry.IsDir() { //append to frontier
				frontier = append(frontier, fullPath)
			} else { //read tags

				// filter out irrelevant files
				loc := strings.LastIndex(entry.Name(), ".")
				extension := entry.Name()[loc+1:]
				if !slices.Contains(PERMITTED_FILE_TYPES, extension) {
					continue; //skip over irrelevant files
				}

				//fmt.Println("Attempting to read: ", fullPath)

				track, err := FFProbeTags(ctx, fullPath)
				if err != nil {
					failed = append(failed, fullPath)
					continue
				}

				tracks = append(tracks, track)

				//fmt.Println("Title: ", track.Title)
				//fmt.Println("Artists: ", track.Artists)
			}
		}
	}

	// now we have all the tracks in a single list
	fmt.Println(tracks)
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


