// Songs are in the form of Artist - Album - Song
//   and as such, albums are Artist - Album to ensure no duplicates

interface Album {
    artists: string[],
    title: string,
    songs: string[],
}


export class AlbumState {
    public albums: Album[] = $state([]);
}







