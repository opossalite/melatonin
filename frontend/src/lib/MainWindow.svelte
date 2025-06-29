<script lang="ts">
    import LeftBar from "./LeftBar/LeftBar.svelte";
    import Content from "./Content/Content.svelte";
    //import { AlbumState } from "./albums.svelte";
    let {albums}: {albums: AlbumState} = $props();

    let barSize = 750;
</script>


<div id="main_window">
    <div id="left_bar" style="width: {barSize}px;"><LeftBar/></div>
    <div id="content" style="left: {barSize}"><Content/></div>
</div>


<style>
#main_window {
    display: flex;
    height: 100%;
}


#content {
    width: 100%;
}
</style>
