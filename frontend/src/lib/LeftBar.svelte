<script lang="ts">
    import type { ProgramState } from "$lib/structures.svelte";
    import { AudioLines } from "lucide-svelte";

    let {program_state}: {program_state: ProgramState} = $props();

    function select(i: number) {
        program_state.selected_album_index = i;
    }
</script>

<div id="wrapper">
    <div id="scroll-area">
        <div id="leftbar">
            {#if program_state.albums.length === 0}
                <p>Reading local files...</p>
            {:else}
                {#each program_state.albums as album, i (album?.title ?? i)}
                    <div class="selected album-box" class:selected={program_state.selected_album_index === i} on:click={() => select(i)}>
                        <div class="album-title">{album.title}</div>
                        <div class="album-artists">{album.artists.join(", ")}</div>
                    </div>
                {/each}
            {/if}

        </div>
    </div>
</div>


<style>
/* outer wrapper stacks the scrolling area + bottom spacer */
#wrapper {
    display: flex;
    flex-direction: column;
    height: 100%;           /* fill the parent's middle row */
    padding: 0.4rem;
    box-sizing: border-box;
    background: var(--base);
}

/* the element that actually scrolls, min-height:0 is CRITICAL so flex children can overflow properly */
#scroll-area {
    flex: 1 1 auto;
    min-height: 0;
    overflow-y: auto;
    overflow-x: hidden;
    overscroll-behavior: contain;
    scrollbar-gutter: stable;
}

/* inner content: same visuals as before but not responsible for scrolling */
#leftbar {
    background: var(--background);
    padding: 0.8rem;
    box-sizing: border-box;
}

/* album item styles unchanged */
.album-box {
    padding: 1rem;
    margin-left: 0.4rem 0;
    margin-right: 0.4rem 0;
    background: var(--background);
    border-radius: 0.5rem;
    width: 100%;
    white-space: nowrap;       /* keep on one line */
    overflow: hidden;          /* cut off the extra text */
    text-overflow: ellipsis;   /* show "..." */
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
    color: var(--main); font-weight: bold; font-size: 1.1rem;
}
.album-artists {
    color: var(--text_dim); font-size: 0.95rem;
}


</style>


