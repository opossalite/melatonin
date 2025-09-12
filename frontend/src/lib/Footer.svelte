<script lang="ts">
    import type { ProgramState } from "$lib/structures.svelte";
    import { SkipBack, SkipForward, CirclePause, CirclePlay, MicVocal, Disc } from "lucide-svelte";

    let {program_state}: {program_state: ProgramState} = $props();

    let isPlaying = $state(false); //if true, pause icon will show
    //note: play and pause seem to be size 24 by default
    //

    function playPauseClick() {
        isPlaying = !isPlaying;
    }

    function skipBack() {

    }

    function skipForward() {

    }
</script>

<div id="footer">
    <div id="track_display" class="group">
        <div>
            Track Title
        </div>
        <div>
            Track Arist
        </div>
    </div>
    <div id="playback_controls" class="group">
        <div id="buttons">
            <SkipBack class="skip" onclick={() => skipBack()}/>
            <CirclePause class="play_pause" onclick={() => playPauseClick()} display={isPlaying ? "block" : "none"}/>
            <CirclePlay class="play_pause" onclick={() => playPauseClick()} display={isPlaying ? "none" : "block"}/>
            <SkipForward class="skip" onclick={() => skipForward()}/>
        </div>
        <div id="progress_info">
            progress bar lol
        </div>
    </div>
    <div id="track_details" class="group">
        <Disc class="now_playing"/>
        <MicVocal class="lyrics"/>
    </div>
</div>

<style>
#footer {
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



#track_display {
    margin-right: auto;
}



#playback_controls {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
}
#buttons {
    display: flex;
    align-items: center;
}
:global(.play_pause) {
    color: white;
    height: 34px; /*should be 24 by default*/
    width: 34px;
    margin: 10px;
}
:global(.play_pause:hover) {
    color: lightgrey;
}

:global(.skip) {
    color: white;
}
:global(.skip:hover) {
    color: lightgrey;
}



#track_details {
    margin-left: auto;
}
:global(.now_playing) {
    color: white;
}
:global(.now_playing:hover) {
    color: lightgrey;
}
:global(.lyrics) {
    color: white;
}
:global(.lyrics:hover) {
    color: lightgrey;
}



</style>
