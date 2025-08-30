<script lang="ts">
    import Header from "$lib/Header/Header.svelte";
    import Footer from "$lib/Footer/Footer.svelte";
    import MainWindow from "$lib/MainWindow.svelte";
    import { onMount } from "svelte";
    import { ProgramState } from "$lib/structures.svelte";



    // establish one singular set of albums that will be maintained by the whole program
    let program_state: ProgramState = $state(new ProgramState);
    onMount(async () => {
        // 

        // retrieve albums, using the settings from before
        const response = await fetch("http://localhost:8800/get_albums");
        const json = await response.json();

        program_state.albums = json.albums;

        //for (var album of albums) {
        //    //program_state..push(album);
        //}
    });
    



</script>

<div id="wrapper">
    <div id="top"><Header/></div>
    <div id="middle"><MainWindow {program_state}/></div>
    <div id="bottom"><Footer {program_state}/></div>
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




