<script lang="ts">
    import { ChevronLeft, ChevronRight, House, Search, ListMusic, Settings } from "lucide-svelte";
</script>

<div id="header">
    <div id="history" class="group">
        <ChevronLeft/>
        <ChevronRight/>
    </div>
    <div id="center_controls" class="group">
        <House/>
        <Search/>
        <ListMusic/>
    </div>
    <div id="settings" class="group">
        <Settings/>

    </div>
</div>

<style>
#header {
    background: grey;
    display: flex;
    height: 100%;
    width: 100%;
    align-items: center;
    box-sizing: border-box;
    padding: 1rem;
    position: relative;
}
.group {
    flex-direction: column;
    gap: 0.5rem;
    background: red;
}



#history {
    margin-right: auto;
}



#center_controls {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
}



#settings {
    align-items: flex-end;
}
</style>
