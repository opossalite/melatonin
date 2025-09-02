# Melatonin
in-progress music player for linux
- playlist structure will be wholely determined by the file structure in storage with 0 in-app organization options
- extremely simple interface focusing on the album art
- smart search options
- artist pages will be generated from files via metadata
- ideally a mobile app to control playback as well
- will easily allow user to open the file manager to reorganize or rename files

### technologies:
go (chi) + svelte + electron

### misc:
lucide (icons)

### notes:
`npm run dev` for react and svelte


### ideas:
- artist pages (contains all the singles they've done and albums)
- search feature will find all songs and artists that fuzzy match
- everything will happen on the / route, with state changing based on the backend's internal knowledge
- in album view, have the number of tracks, the total length, and length remaining
- album view should have album art that can be seen even when listening to something else
- history will be tracked as frames in the backend that can easily be brought back
- a bottom right button will open up the user's default file manager for easy file management
- vim-compatible keys that will remain mostly hidden when using a mouse, like searchbar not being key-selectable but the search icon being selectable
    - will be able to exit vim mode by pressing esc but can enter it by pressing any vim key
    - figure out how to right click and play and such the songs, maybe enter for play and ctrl+y for right clicking?
    - press v to select or deselect a track, and shift+v selects all between last selected and currently hovering
- customizable color theme (everything will remain the same but this one color)
- main color orange, and then vim mode a light teal?


