<script lang="ts">
    import Header from "$lib/Header/Header.svelte";
    import Footer from "$lib/Footer/Footer.svelte";
    import MainWindow from "$lib/MainWindow.svelte";
    import { onMount } from "svelte";
    import { AlbumState } from "$lib/albums.svelte";
    import { Settings } from "$lib/settings.svelte";



    // establish one singular set of albums that will be maintained by the whole program
    let settings: Settings;
    let albums: AlbumState = new AlbumState;
    onMount(async () => {
        // 

        // retrieve albums, using the settings from before
        const response = await fetch("http://localhost:8800/get_albums");
        const json = await response.json();

        for (var album of json.albums) {
            albums.albums.push(album);
        }
    });
    



</script>

<div id="wrapper">
    <div id="top"><Header/></div>
    <div id="middle"><MainWindow {albums}/></div>
    <div id="bottom"><Footer/></div>
</div>





<style>


:global(html,body),#wrapper {
    height:100%;
    padding:0;
    margin:0;
}
#wrapper {
    position: relative;
}

#top, #middle, #bottom {
    position:absolute;
}

#top {
    height:60px;
    width:100%;
    background:grey;
}
#middle {
    top:60px;
    bottom:80px;
    width:100%;
    background:black;
    color:white;
}
#bottom {
    bottom:0;
    height:80px;
    width:100%;
}

</style>




