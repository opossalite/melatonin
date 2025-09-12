<script lang="ts">
    import { browser } from "$app/environment";
    import Header from "$lib/Header.svelte";
    import Footer from "$lib/Footer.svelte";
    import MainWindow from "$lib/MainWindow.svelte";
    import { onMount } from "svelte";
    import { ProgramState, type Album } from "$lib/structures.svelte";



    interface AlbumsResponse {
        albums: Album[];
    }

    // establish one singular set of albums that will be maintained by the whole program
    let program_state: ProgramState = $state(new ProgramState);
    onMount(async () => {
        //// retrieve albums, using the settings from before
        const response = await fetch("http://localhost:8800/get_albums", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ folders: program_state.folders }),
        });
        const json: AlbumsResponse = await response.json();
        program_state.albums = json.albums;
    });

    $effect(() => {
        if (!browser) { //hmmmmm
            return;
        }
        const root = document.documentElement;
        root.style.setProperty("--main", program_state.tupleToHex(program_state.main_color));
        root.style.setProperty("--base", program_state.tupleToHex(program_state.base_color));
        root.style.setProperty("--background", program_state.tupleToHex(program_state.background_color));
        root.style.setProperty("--highlight", program_state.tupleToHex(program_state.highlight_color));
        root.style.setProperty("--selected", program_state.tupleToHex(program_state.selected_color));
        root.style.setProperty("--selected_highlight", program_state.tupleToHex(program_state.selected_highlight_color));
        root.style.setProperty("--text", program_state.tupleToHex(program_state.text_color));
        root.style.setProperty("--text_dim", program_state.tupleToHex(program_state.text_color_dim));
    });
    



</script>

<div id="wrapper">
    <div id="top"><Header/></div>
    <div id="middle"><MainWindow {program_state}/></div>
    <div id="bottom"><Footer {program_state}/></div>
</div>





<style>


:global(html,body), #wrapper {
    height: 100%;
    padding: 0;
    margin: 0;
}

#wrapper {
    display: grid;
    grid-template-rows: 60px 1fr 80px; /* top / middle / bottom */
    min-height: 100vh;
}

/* place children into rows */
#top { grid-row: 1; height: 60px; }
#middle { grid-row: 2; overflow: hidden; } /* middle fills remaining space */
#bottom { grid-row: 3; height: 80px; }

</style>




