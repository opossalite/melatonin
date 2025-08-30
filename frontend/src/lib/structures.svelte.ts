// Songs are in the form of Artist - Album - Song
//   and as such, albums are Artist - Album to ensure no duplicates

//interface Album {
//    artists: string[],
//    title: string,
//    songs: string[],
//}
//
//
//export class AlbumState {
//    public albums: Album[] = $state([]);
//}

interface Track {
    title: string,
    artists: string[],
    album: string,
    album_artists: string[],
    year: string,
    track_no: string,
    track_count: string,
    cd_no: string,
    cd_count: string,
    //eventually add cover as well

}



export class ProgramState {
    // runtime
    public albums: Track[][] = $state([]);

    // program settings
    public folders: string[] = []
    public main_color: [number, number, number] = [0, 0, 0]
}







