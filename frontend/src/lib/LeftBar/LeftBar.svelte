<script lang="ts">
    import type { ProgramState } from "$lib/structures.svelte";
    import { AudioLines } from "lucide-svelte";

    let {program_state}: {program_state: ProgramState} = $props();
</script>

<div id="leftbar">
    left bar
    <AudioLines/>

    {#each program_state.albums as album ((album[0].artists, album[0].title))}
        <div class="album-box">
            <div class="album-title">{album[0].title}</div>
            <div class="album-artists">{album[0].artists.join(", ")}</div>
        </div>
    {/each}
</div>


<style>
#leftbar {
    /*background: lime;*/
    background: #121212;
    height: 100%;
}

.album-box {
    padding: 1rem;
    margin: 0.5rem 0;
    background-color: #121212;
}
.album-box:hover {
    background-color: #1f1f1f;
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
