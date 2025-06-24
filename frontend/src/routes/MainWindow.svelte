<script lang="ts">
    import LeftBar from "./LeftBar.svelte";
    import {createState, NathanState} from "./state.svelte";

    let {name}: {name: string} = $props();

    const myState = createState();
    const myStateTwo = new NathanState();
</script>


<div id="main_window">
    main window
    <h1>
        {name ? name : "User"}'s Form
    </h1>
    <button onclick={myState.up}>{myState.value}</button>
    <button onclick={() => myStateTwo.up()}>{myStateTwo.value}</button>
</div>


<style>
</style>
