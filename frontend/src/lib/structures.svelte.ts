


export interface Track {
    title: string;
    artists: string[];
    album: string;
    album_artists: string[];
    year: number;
    track: number;
    track_total: number;
    disc: number;
    disc_total: number;
    duration: number;
    bitrate: number;
    path: string;
}



export interface Disc {
    tracks: Track[];
}



export interface Album {
    artists: string[];
    title: string;
    year: number;
    discs: Disc[];
}


export class ProgramState {
    // runtime
    public albums: Album[] = $state([]);
    public selected_album_index: number | null = $state(null);

    // program settings
    public folders: string[] = [
        "~/Music",
    ]

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







