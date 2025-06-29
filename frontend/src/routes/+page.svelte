<script lang="ts">
    import { fly } from "svelte/transition";
    import Header from "$lib/Header/Header.svelte";
    import Footer from "$lib/Footer/Footer.svelte";
    import MainWindow from "$lib/MainWindow.svelte";
    import { AlbumState }  from "$lib/albums.svelte.ts";

    async function roll() {
        const response = await fetch("http://localhost:8800/roll");
        const json = await response.json();
        //roll_result = json.value;
    }

    let albums = new AlbumState();
    



</script>

<div id="wrapper">
    <div id="top"><Header name="a"/></div>
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




