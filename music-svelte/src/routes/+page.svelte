<script lang="ts">
    import Header from "./Header.svelte";

    let formState = $state({
        answers: {},
        step: 0,
        error: "",
    });

    const QUESTIONS = [
        {
            question: "What is your name?",
            id: "name",
            type: "text",
        },
        {
            question: "What's your birthday?",
            id: "birthday",
            type: "date",
        },
        {
            question: "What's your favorite color?",
            id: "color",
            type: "color",
        },
    ];

    function nextStep(id: string) {
        if (formState.answers[id]) {
            formState.step += 1;
            formState.error = "";
        } else {
            formState.error = "Please fill out the form input."
        }
    }
</script>





<Header name={formState.answers.name}/>

<main>
    {#if formState.step >= QUESTIONS.length}
        <p>Thank you!</p>
    {:else}
        <p>Step: {formState.step + 1}</p>
    {/if}

    <!-- {#each QUESTIONS as question(question.id)} -->
    {#each QUESTIONS as question, i (question.id)}
        {#if (formState.step === i)}
            {@render formStep(question)}
        {/if}
    {/each}


    {#if formState.error}
        <p class="error">{formState.error}</p>
    {/if}
</main>

<!-- {JSON.stringify(formState)} -->


{#snippet formStep({question, id, type}: {question: string, id: string, type: string})}
    <article>
        <div>
            <label for={id}>{question}</label>
            <input {type} {id} bind:value={formState.answers[id]}>
        </div>
        <button onclick={() => nextStep(id)}>Next</button>
    </article>
{/snippet}




<style>
    .error {
        color: red;
    }
</style>




