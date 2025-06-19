<script lang="ts">
    import Header from "./Header.svelte";

    let formState = $state({
        name: "",
        birthday: "",
        step: 0,
        error: "",
    });

</script>

<Header name={formState.name}/>

<main>
    <p>Step: {formState.step}</p>
    {@render formStep({question: "What's your name?", id: "name", type: "text"})}
    {#if formState.error}
        <p class="error">{formState.error}</p>
    {/if}
</main>


{#snippet formStep({question, id, type}: {question: string, id: string, type: string})}
    <article>
        <div>
            <label for={id}>{question}</label>
            <input {type} {id} bind:value={formState[id]}>
        </div>
    </article>
{/snippet}



<style>
    .error {
        color: red;
    }
</style>




