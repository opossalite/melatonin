# Psilocybin
in-progress music player for linux
- playlist structure will be wholely determined by the file structure in storage with 0 in-app organization options
- extremely simple interface focusing on the album art
- smart search options
- artist pages will be generated from files via metadata
- ideally a mobile app to control playback as well
- will easily allow user to open the file manager to reorganize or rename files

### technologies:
rust (rodio, rocket) + svelte + electron

### misc:
lucide (icons)

### notes:
`npm run dev` for react and svelte


### ideas:
- artist pages (contains all the singles they've done and albums)
- search feature will find all songs and artists that fuzzy match
- everything will happen on the / route, with state changing based on rust's internal knowledge
- in tracklist view, have the number of tracks, the total length, and length remaining
- tracklist view should have album art so that it can be seen even when listening to something else
- history will be tracked as blocks in rust that can easily be brought back
- bottom right button will open up their default file manager
- vim-compatible keys that will remain mostly hidden when using a mouse, like searchbar not being key-selectable but the search icon being selectable
    - will be able to exit vim mode by pressing esc but can enter it by pressing any vim key
    - figure out how to right click and play and such the songs, maybe enter for play and ctrl+y for right clicking?
- customizable color theme (everything will remain the same but this one color)
- press v to select or deselect a track, and shift+v selects all between last selected and currently hovering
- main color orange, and then vim mode a light teal?


