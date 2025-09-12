<script lang="ts">
    let {name}: {name: string} = $props();
    let display = $state("Home")

    const myState = createState();
    const myStateTwo = new NathanState();
    import {createState, NathanState} from "./state.svelte";
</script>

<div id="content">
    {#if display === "Home"}
        Home Page
        <h1>
            {name ? name : "User"}'s Form
        </h1>
        <button onclick={myState.up}>{myState.value}</button>
        <button onclick={() => myStateTwo.up()}>{myStateTwo.value}</button>
    {:else if display === "Queue"}
        <div>
            Queue
        </div>
    {:else if display === "Album"}
        <div>
            Album
        </div>
    {:else if display === "Singles"}
        <div>
            Singles
        </div>
    {:else if display === "Search"}
        <div>
            Search
        </div>
    {:else if display === "Artist"}
        <div>
            Artist
        </div>
    {:else if display === "Now Playing"}
        <div>
            Now Playing
        </div>
    {:else if display === "Lyrics"}
        <div>
            Lyrics
        </div>
    {:else if display === "Settings"}
        <div>
            Settings
        </div>
    {:else}
        <div>
            If you ever see this, please open a GitHub issue, thank you.<br>
            Include this: display = {display}
        </div>
    {/if}
</div>


<style>
#content {
    background: teal;
    height: 100%;
    width: 100%;
}
</style>
