<script lang="ts">
    let {name, fake_name}: {name: string, fake_name?: string} = $props();
</script>

<h1>Hello {name}</h1>
<h2>{name.replaceAll("a", "v")}</h2>
<h2>{fake_name}</h2>


<style>
</style>

