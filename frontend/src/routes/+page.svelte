<script lang="ts">
    import { fly } from "svelte/transition";
    import Header from "./Header.svelte";

    




    let roll_result = $state("-");
    async function roll() {
        //const response = await fetch("http://localhost:8800/roll", {
        //    method: "POST",
        //    headers: {
        //        "Content-Type": "application/json"
        //    },
        //    body: JSON.stringify({ message: "roll" })
        //});
        const response = await fetch("http://localhost:8800/roll");
        //number = response.json(); //might have to await
        const json = await response.json();
        roll_result = json.value;
    }

    let word = $state("");
    let word_upper = $state("");
    async function upper() {
        const response = await fetch("http://localhost:8800/upper", {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({word}),
        });
        const json = await response.json();
        word_upper = json.word_upper;
    }
    $effect(() => {
        upper();
    })



    let formState = $state({
        answers: {},
        step: 0,
        error: "",
    });


    $inspect(formState.step);



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

    $effect(() => {
        console.log("on mounted");
        return () => {
            console.log("on unmounted");
        };
    });

    $effect(() => {
        // this will run whenever formState.step has changed, and upon initialization
        console.log("formState", formState.step);

        //before effect re-runs
        return () => {
            console.log("before formState returns", formState.step)
        };
    });

</script>





<Header name={formState.answers.name}/>
<br>
<div>
    Roll some dice!
    <button onclick={() => roll()}>{roll_result}</button>
</div>
<br>
<div>
    <input type="text" bind:value={word}>
    <br>
    -> {word_upper}

</div>
<main>
    {#if formState.step >= QUESTIONS.length}
        <p>Thank you!</p>
    {:else}
        <p>Step: {formState.step + 1}</p>
    {/if}

    <!-- {#each QUESTIONS as question(question.id)} -->
    {#each QUESTIONS as question, i (question.id)}
        {#if (formState.step === i)}
            <div in:fly={{x: 200, duration: 200, opacity: 0, delay: 200}} out:fly={{x: -200, duration: 200, opacity: 0}}>
                {@render formStep(question)}
            </div>
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

    .horiz {
        display: inline-block;
    }
</style>




