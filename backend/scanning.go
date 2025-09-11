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
	"sync"
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
				Track:    	  1,
				TrackTotal:   10,
				Disc:         1,
				DiscTotal:      1,
			},
			{
				Title:        "Song B",
				Artists:      []string{"Artist One"},
				Album:        "Greatest Hits",
				AlbumArtists: []string{"Artist One"},
				Year:         2024,
				Track:      	2,
				TrackTotal:   10,
				Disc:         1,
				DiscTotal:      1,
			},
		},
		{
			{
				Title:        "Another Song",
				Artists:      []string{"Different Band"},
				Album:        "Live at Home",
				AlbumArtists: []string{"Different Band"},
				Year:         2019,
				Track:      	1,
				TrackTotal:   8,
				Disc:         1,
				DiscTotal:      1,
			},
		},
	};

}



// Run ffprobe in the command line for a given file path and return a Track
func FFProbeTags(ctx context.Context, path string, solo bool) (Track, error) {
    if _, ok := ctx.Deadline(); !ok {
        var cancel context.CancelFunc
        ctx, cancel = context.WithTimeout(ctx, 10*time.Second)
        defer cancel()
    }

	show := "format=duration,bit_rate:" +
        "format_tags=title,artist,album,album_artist,date,track,tracktotal,disc,disctotal"

    cmd := exec.CommandContext(
        ctx, "ffprobe",
        "-v", "error",
        "-of", "default=nw=1:nk=0", // no wrappers, show keys
        "-show_entries", show,
        path,
    )

	// set up an stdout and an stderr
    var stdout, stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    if err := cmd.Run(); err != nil {
        // if ffprobe printed something, surface it in the error
        msg := strings.TrimSpace(stderr.String())
        if msg == "" {
            msg = err.Error()
        }
        return Track{}, errors.New("ffprobe: " + msg)
    }

    sc := bufio.NewScanner(bytes.NewReader(stdout.Bytes()))
	var track Track
    for sc.Scan() {

		// set up a pipeline to extract valuable info
        line := strings.ToLower(strings.TrimSpace(sc.Text()))
        if line == "" {
			continue
        }
		loc := strings.Index(line, "=")
		key := line[:loc]
		value := line[loc+1:]

		//fmt.Println("key: ", key)
		//fmt.Println("value: ", value)

		switch key {
		case "duration": //no way this breaks
			num, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: Duration] %w", err)
			}
			track.Duration = num

		case "bit_rate": //no way this breaks
			num, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: Bitrate] %w", err)
			}
			track.Bitrate = num

		case "tag:title":
			track.Title = value

		case "tag:artist":
			track.Artists = strings.Split(value, ";")

		case "tag:album":
			if solo { continue }
			track.Album = value

		case "tag:album_artist":
			if solo { continue }
			track.AlbumArtists = strings.Split(value, ";")

		case "tag:date":
			num, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: Date] %w", err)
			}
			track.Year = num

		case "tag:track":
			if solo { continue }
			num, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: Track] %w", err)
			}
			track.Track = num

		case "tag:tracktotal":
			if solo { continue }
			num, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: TrackTotal] %w", err)
			}
			track.TrackTotal = num

		case "tag:disc":
			if solo { continue }
			num, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: Disc] %w", err)
			}
			track.Disc = num

		case "tag:disctotal":
			if solo { continue }
			num, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return Track{}, fmt.Errorf("[TAG: DiscTotal] %w", err)
			}
			track.DiscTotal = num
		}
	}

    if err := sc.Err(); err != nil {
        return Track{}, err
    }

	track.Path = path //add the path now
	if solo {
		track.Album = "SOLO"
		track.AlbumArtists = []string{"SOLO"}
	}

    return track, nil
}




//func readAlbums(folders []string) [][]Track {
func readTracks(folders []string, exclude_folders []string) []Track {
	fmt.Println("Reading local files...")
	frontier := expandAll(folders)
	exclude := expandAll(exclude_folders)

	tracks := []Track{} //will be sorted later
	mu := sync.Mutex{}
	//failed := []string{}

	ctx := context.Background() //for commandline use

	wg := sync.WaitGroup{}

	for len(frontier) > 0 {
		path := frontier[len(frontier)-1]
		frontier = frontier[:len(frontier)-1]

		entries, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("Error:", err)
			return []Track{}
		}

		for _, entry := range entries {
			fullPath := filepath.Join(path, entry.Name())
			if slices.Contains(exclude, fullPath) {
				continue;
			}

			var solo bool
			if strings.Index(strings.ToLower(path), ",solos/") > 0 {
				solo = true
			} else {
				solo = false
			}

			if entry.IsDir() { //append to frontier
				frontier = append(frontier, fullPath)
			} else { //read tags concurrently

				wg.Add(1)
				go func() {
					defer wg.Done()

					// filter out irrelevant files
					loc := strings.LastIndex(entry.Name(), ".")
					extension := entry.Name()[loc+1:]
					if !slices.Contains(PERMITTED_FILE_TYPES, extension) {
						return; //skip over irrelevant files
					}

					track, err := FFProbeTags(ctx, fullPath, solo)
					if err != nil {
						fmt.Println("FAILED TO READ:", fullPath, "-----", err)
						//failed = append(failed, fullPath)
						return
					}

					// now just ensure we have all the tags we need
					conditions := []bool{
						track.Title == "",
						track.Artists == nil,
						track.Year == 0,
						!solo && track.Album == "",
						!solo && track.AlbumArtists == nil,
						!solo && track.Track == 0,
						!solo && track.TrackTotal == 0,
						!solo && track.Disc == 0,
						!solo && track.DiscTotal == 0,
					}
					messages := []string{
						"Title",
						"Artist",
						"Year",
						"Album",
						"Album_Artists",
						"Track",
						"TrackTotal",
						"Disc",
						"DiscTotal",
					}

					is_missing_tags := false
					for i := range len(conditions) {
						if conditions[i] {
							is_missing_tags = true
							fmt.Printf("ERROR: MISSING TAG '%s' for '%s'\n", messages[i], fullPath)
						}
					}
					if is_missing_tags {
						return
					}

					mu.Lock()
					tracks = append(tracks, track)
					mu.Unlock()
				}()
			}
		}
	}

	wg.Wait()

	//fmt.Println(tracks)

	return tracks
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


