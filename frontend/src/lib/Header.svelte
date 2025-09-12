<script lang="ts">
    import { ChevronLeft, ChevronRight, House, Search, AlignJustify, Settings } from "lucide-svelte";
</script>

<div id="header">
    <div id="history" class="group">
        <ChevronLeft class="back"/>
        <ChevronRight class="forward"/>
    </div>
    <div id="center_controls" class="group">
        <House class="home"/>
        <Search class="search_icon"/>
        <AlignJustify class="playlist_view"/>
    </div>
    <div id="settings" class="group">
        <Settings class="settings_icon"/>

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
}



#history {
    margin-right: auto;
}
:global(.back) {
    color: white;
}
:global(.back:hover) {
    color: lightgrey;
}
:global(.forward) {
    color: white;
}
:global(.forward:hover) {
    color: lightgrey;
}



#center_controls {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
}
:global(.home) {
    color: white;
}
:global(.home:hover) {
    color: lightgrey;
}
:global(.search_icon) {
    color: white;
}
:global(.search_icon:hover) {
    color: lightgrey;
}
:global(.playlist_view) {
    color: white;
}
:global(.playlist_view:hover) {
    color: lightgrey;
}



#settings {
    align-items: flex-end;
}
:global(.settings_icon) {
    color: white;
}
:global(.settings_icon:hover) {
    color: lightgrey;
}
</style>
