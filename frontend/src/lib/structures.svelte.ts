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
    public selected_album_index: number | null = $state(null);

    // program settings
    public folders: string[] = []

    // colors
    public main_color: [number, number, number] = [215, 142, 30] //accent
    public base_color: [number, number, number] = [0, 0, 0] //used for borders and the top and bottom panels
    public background_color: [number, number, number] = [18, 18, 18] //background of each panel
    public highlight_color: [number, number, number] = [31, 31, 31] //when anything background_color is highlighted
    public selected_color: [number, number, number] = [42, 42, 42] //when anything background_color is highlighted
    public selected_highlight_color: [number, number, number] = [72, 72, 72] //when anything background_color is highlighted
    public text_color: [number, number, number] = [255, 255, 255];
    public text_color_dim: [number, number, number] = [187, 187, 187];

    // converts an RGB tuple to a hex color string (e.g., [255, 51, 0] → "#ff3300")
    public rgbToHex(r: number, g: number, b: number): string {
        const toHex = (n: number) => Math.max(0, Math.min(255, n)).toString(16).padStart(2, '0');
        return `#${toHex(r)}${toHex(g)}${toHex(b)}`;
    }

    // convenience method to convert a tuple
    public tupleToHex(rgb: [number, number, number]): string {
        return this.rgbToHex(rgb[0], rgb[1], rgb[2]);
    }
}







