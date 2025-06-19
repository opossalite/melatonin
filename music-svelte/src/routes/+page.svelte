<script lang="ts">
    import Header from "./Header.svelte";

    let name = $state("Nathan");
    let status: "OPEN" | "CLOSED" = $state("OPEN");
    let full_name = $derived(name + " " + "Bacon")

    function onclick() {
        status = status === "OPEN" ? "CLOSED" : "OPEN";
    }
</script>

<Header {name} fake_name="Mason"/>
<h2>{full_name}</h2>
<input type="text" bind:value={name}/>
<p>The store is now {status}</p>
<button {onclick}>Toggle Status</button>


