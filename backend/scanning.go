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
	"sort"
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

	//fmt.Println("ALBUM NAME:")
	//fmt.Println(track.AlbumArtists, track.Album)
	//if track.Album == "ice queen" {
	//	fmt.Println("ICE QUEEN")
	//	fmt.Println(solo)
	//	fmt.Println(track.Path)
	//	fmt.Println(strings.Contains(strings.ToLower(track.Path), ",solos/"))
	//}

    return track, nil
}




func readTracks(folders []string, exclude_folders []string) []Track {
	fmt.Println("Reading local files...")
	frontier := expandAll(folders)
	exclude := expandAll(exclude_folders)

	tracks := []Track{} //will be sorted later
	//failed := []string{}
	mu := sync.Mutex{}

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

					// establish solo variable
					var solo bool
					if strings.Contains(strings.ToLower(fullPath), ",solos/") {
						solo = true
					} else {
						solo = false
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
	return tracks
}



type AlbumKey struct {
	ArtistsConcat string
	Album string
}
type AlbumConstructor struct {
	Artists []string
	Title string
	Year uint64
	Tracks []Track
}
// Sort a list of tracks into a list of albums
func sortTracks(tracks []Track) ([]Album) {

	// first create albums and assign tracks to them
	album_c_map := make(map[AlbumKey]*AlbumConstructor)
	
	for i := range len(tracks) {
		//artists := tracks[i].AlbumArtists
		track := tracks[i]
		artists := make([]string, 0, len(track.AlbumArtists))
		copy(artists, track.AlbumArtists)
		sort.Strings(artists)
		artists_concat := strings.Join(artists, ";")
		key := AlbumKey {artists_concat, track.Album}

		album, ok := album_c_map[key]
		if !ok {
			album = &AlbumConstructor{
				Artists: track.AlbumArtists,
				Title: track.Album,
				Year: 0,
				Tracks: []Track{track},

			}
			album_c_map[key] = album
		} else {
			album.Tracks = append(album.Tracks, track)
		}
	}

	// extract the album constructors from the map into a list instead
	album_cs := make([]*AlbumConstructor, 0, len(album_c_map))
	for _, album_c := range album_c_map {
		album_cs = append(album_cs, album_c)
	}

	//for _, x := range album_cs {
	//	fmt.Println(*x)
	//}
	//fmt.Println("ugh")

	// iterate through all albums, processing their discs and tracks
	album_collection := []Album{}
	for _, album_c := range album_cs {
		//fmt.Println("NOW PROCESSING:")
		//fmt.Println(album_c)
		//fmt.Println("UGH")

		tracks := album_c.Tracks

		// handle solos differently, without much processing at all
		if album_c.Title == "SOLO" && slices.Equal(album_c.Artists, []string{"SOLO"}) {
			new_album := Album{
				Artists: album_c.Artists,
				Title: album_c.Title,
				Year: album_c.Year,
				Discs: []Disc{Disc{Tracks: album_c.Tracks}},
			}
			album_collection = append(album_collection, new_album)
			continue
		}

		// used throughout this loop
		passed := true

		// sort the tracks by track then disc numbers
		sort.Slice(tracks, func(i, j int) bool {
			return tracks[i].Track < tracks[j].Track
		})
		sort.SliceStable(tracks, func(i, j int) bool {
			return tracks[i].Disc < tracks[j].Disc
		})

		//fmt.Println("AFTER SORTING")
		//fmt.Println(album_c)
		//fmt.Println("URGH")

		// establish a disc map and place the tracks into the right disc
		disc_map := map[uint64]*Disc{}
		disc_total := album_c.Tracks[0].DiscTotal
		for i, track := range tracks {
			// throw tracks into the right disc
			if track.DiscTotal != disc_total {
				fmt.Println("FAILED ALBUM [DISCTOTAL MISMATCH BETWEEN INDEX i AND INDEX 0]:", album_c.Artists, album_c.Title, track.Path, album_c.Tracks[0].Path)
				fmt.Println("--INDEX", i)
				fmt.Println("--INDEX i DISCTOTAL", track.DiscTotal)
				fmt.Println("--INDEX 0 DISCTOTAL", album_c.Tracks[0].DiscTotal)
				passed = false
				break
			}

			// if the disc already exists, add to it, otherwise create it
			disc, ok := disc_map[track.Disc]
			if !ok {
				disc = &Disc{
					Tracks: []Track{track},
				}
				disc_map[track.Disc] = disc
			} else {
				disc.Tracks = append(disc.Tracks, track)
			}
		}
		if !passed {
			continue
		}

		// now convert discs into a list and ensure their numerations make sense
		discs := []Disc{}
		for i := range uint64(len(disc_map)) {
			i += 1
			disc, ok := disc_map[i]
			if !ok {
				fmt.Println("FAILED ALBUM [ALBUM DISC NUMBERS NOT INCREMENTAL]:", album_c.Artists, album_c.Title)
				fmt.Println("--DISC_MAP LENGTH", len(disc_map))
				fmt.Println("--ATTEMPTED TO GRAB", i)
				passed = false
				break
			} else {
				discs = append(discs, *disc)
			}
		}
		if !passed || uint64(len(discs)) != disc_total || uint64(len(disc_map)) != disc_total {
			continue
		}

		// process the track numbers and establish years, now that the number of discs is verified and their creation contains the right tracks
		year := uint64(0)
		for _, disc := range discs {
			for i, track := range disc.Tracks {
				i += 1
				if track.Track != uint64(i) || track.TrackTotal != uint64(len(disc.Tracks)) {
					fmt.Println("FAILED ALBUM [MISMATCHED TRACKTOTAL OR TRACK TAG]:", album_c.Artists, album_c.Title)
					passed = false
					break
				}
				if track.Year > year {
					year = track.Year
				}
			}
		}
		if !passed {
			continue
		}

		// now that track numbers are verified, compile it all into a complete album
		new_album := Album{
			Artists: album_c.Artists,
			Title: album_c.Title,
			Year: year,
			Discs: discs,
		}
		album_collection = append(album_collection, new_album)
	}

	return album_collection
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


