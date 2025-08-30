<script lang="ts">
    import LeftBar from "./LeftBar/LeftBar.svelte";
    import Content from "./Content/Content.svelte";
    import { ProgramState } from "./structures.svelte";
    let {program_state}: {program_state: ProgramState} = $props();

    let barSize = 750;
</script>


<div id="main_window">
    <div id="left_bar" style="width: {barSize}px;"><LeftBar {program_state}/></div>
    <div id="content" style="left: {barSize}"><Content name="Opossalite"/></div>
</div>


<style>
#main_window {
    display: flex;
    height: 100%;
}


#content {
    width: 100%;
}
</style>
