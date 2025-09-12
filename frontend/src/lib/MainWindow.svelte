<script lang="ts">
    import LeftBar from "./LeftBar.svelte";
    import Content from "./Content.svelte";
    import { ProgramState } from "./structures.svelte";
    let {program_state}: {program_state: ProgramState} = $props();

    let barSize = $state(750);

    let isDragging = false;

    function startDrag(event: MouseEvent) {
        isDragging = true;
        document.addEventListener("mousemove", onDrag);
        document.addEventListener("mouseup", stopDrag);
        event.preventDefault();
    }

    function onDrag(event: MouseEvent) {
        if (!isDragging) return;
        const newWidth = event.clientX;
        if (newWidth > 200 && newWidth < 1200) { // optional min/max width
            barSize = newWidth;
        }
    }

    function stopDrag() {
        isDragging = false;
        document.removeEventListener("mousemove", onDrag);
        document.removeEventListener("mouseup", stopDrag);
    }
</script>


<div id="main_window">
    <div id="left_bar" style="width: {barSize}px;">
        <LeftBar {program_state}/>
    </div>
    <div id="drag_handle" onmousedown={startDrag}></div>
    <div id="content" style="left: {barSize}">
        <Content name="Opossalite"/>
    </div>
</div>


<style>
#main_window {
    display: flex;
    height: 100%;
}


#content {
    width: 100%;
}


#drag_handle {
    top: 0;
    right: 0;
    width: 10px;              /* thickness of the handle */
    height: 100%;
    cursor: ew-resize;        /* horizontal resize cursor */
    z-index: 10;
    transition: background 0.2s;
    background: var(--base);
}

#drag_handle:hover {
    /*background: rgba(255,255,255,0.2); */
}

</style>
