package main



// Outlines the types of history frames that exist for the main content section
type HistoryType int
const (
	Home HistoryType = iota //customized to whatever content should be on the home
	AlbumView //will display an album or single, whether playing already or not
	ArtistPage //display an artist's page with all their content
	Search //search results
	NowPlaying //album art or whatever other content for now playing
	Lyrics //lyrics sourced from internet
	Settings //settings
)



// Contains the information required to display content on the main content section
type HistoryFrame struct {
	Type   HistoryType

	// only some of these will be filled depending on Type, specified below
	Artist string `json:"artist,omitempty"` //AlbumView, ArtistPage, NowPlaying, Lyrics
	Album string `json:"album,omitempty"` //AlbumView, NowPlaying, Lyrics
	Query string `json:"query,omitempty"` //Search
	Track string `json:"track,omitempty"` //NowPlaying, Lyrics
}



