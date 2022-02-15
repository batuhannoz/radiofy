<template>
  <div class="playback">
    <button @click="Play">Play</button>
  </div>
</template>

<script>
import {play} from "../services/SpotifyPlayer.js"

export default {
  methods: {
    Play() {
      play()
    }
  }
}
</script>




<style>
.playback{
  position: fixed;
  bottom: 0;
  width: 60%;
  height: 84px;
  background-color: grey;
}
</style>