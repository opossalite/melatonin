<script lang="ts">
    import type { Snippet } from "svelte";

    let {name, children, secondChild}: {name: string, children: Snippet, secondChild: Snippet} = $props();
</script>

<div>
    <h1>
        {name ? name : "User"}'s Form
    </h1>

    {@render children()}
    <h3>{@render secondChild("Scott")}</h3>
</div>


