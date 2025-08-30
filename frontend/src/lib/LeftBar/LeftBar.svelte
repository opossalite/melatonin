<script lang="ts">
    import type { ProgramState } from "$lib/structures.svelte";
    import { AudioLines } from "lucide-svelte";

    let {program_state}: {program_state: ProgramState} = $props();

    function select(i: number) {
        program_state.selected_album_index = i;
    }
</script>

<div id="wrapper">
    <div id="leftbar">
        left bar
        <AudioLines/>

        {#each program_state.albums as album, i (album?.[0]?.title ?? i)}
            <div
                class="selected album-box"
                class:selected={program_state.selected_album_index === i}
                onclick={() => select(i)}
                onkeydown={() => {}}
            >
                <div class="album-title">{album[0].title}</div>
                <div class="album-artists">{album[0].artists.join(", ")}</div>
            </div>
        {/each}
    </div>
</div>


<style>
#wrapper {
    padding: 0.4rem;
    background: var(--base);
    height: 100%;
}


#leftbar {
    background: var(--background);
    padding: 0.8rem;
    height: 100%;
    border-radius: 0.5rem;
}

.album-box {
    padding: 1rem;
    margin-left: 0.4rem 0;
    margin-right: 0.4rem 0;
    background: var(--background);
    border-radius: 0.5rem;
}
.album-box:hover {
    background: var(--highlight);
}
.selected {
    background: var(--selected);
}
.selected:hover {
    background: var(--selected_highlight);
}
.album-title {
    color: var(--main);
    font-weight: bold;
    font-size: 1.1rem;
}
.album-artists {
    color: var(--text_dim);
    font-size: 0.95rem;
}
</style>


