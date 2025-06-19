<script lang="ts">
    let {name, fake_name}: {name: string, fake_name?: string} = $props();
</script>

<div>
    <h1>
        {name ? name : "User"}'s Form
    </h1>
</div>


