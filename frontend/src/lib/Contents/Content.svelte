<script lang="ts">
    let {name}: {name: string} = $props();

    const myState = createState();
    const myStateTwo = new NathanState();
    import {createState, NathanState} from "./state.svelte";
</script>

<div>
    content
    <h1>
        {name ? name : "User"}'s Form
    </h1>
    <button onclick={myState.up}>{myState.value}</button>
    <button onclick={() => myStateTwo.up()}>{myStateTwo.value}</button>
</div>

