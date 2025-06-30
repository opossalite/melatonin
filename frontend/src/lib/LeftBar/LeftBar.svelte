<script lang="ts">
    import type { AlbumState } from "$lib/albums.svelte";
    import { AudioLines } from "lucide-svelte";

    let {albums}: {albums: AlbumState} = $props();
</script>

<div id="leftbar">
    left bar
    <AudioLines/>

    {#each albums.albums as album ((album.artists, album.title))}
        <div class="album-box">
            <div class="album-title">{album.title}</div>
            <div class="album-artists">{album.artists.join(", ")}</div>
        </div>
    {/each}
</div>


<style>
#leftbar {
    background: lime;
    height: 100%;
}

.album-box {
    border: 1px solid #ccc;
    padding: 1rem;
    margin: 0.5rem 0;
    border-radius: 8px;
    background-color: #121212;
}
.album-title {
    color: #d78e1e;
    font-weight: bold;
    font-size: 1.1rem;
}
.album-artists {
    color: #bbbbbb;
    font-size: 0.95rem;
}
</style>
