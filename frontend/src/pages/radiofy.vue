<template>
  <div style="height: 100%">
    <div style="height: 50%">
    <va-icon @click="skipToPrevious" color="#000000" size="3rem" name="skip_previous" />
    <va-icon @click="PlayPlayback" v-if="isPlay" color="#000000" size="3rem" name="play_circle" />
    <va-icon @click="PausePlayback" v-if="!isPlay" color="#000000" size="3rem" name="pause_circlen" />
    <va-icon @click="skipToNext" color="#000000" size="3rem" name="skip_next" />
    </div>

    <div class="volume" style="height: 50%"><va-slider v-model="watchVolume" color="black" ></va-slider></div>
  </div>
</template>

<script>
import SpotifyWebApi from "spotify-web-api-js";
const spotifyApi = new SpotifyWebApi();
import VueCookies from 'vue-cookies';

if (VueCookies.get("access_token")) {
  console.log(VueCookies.get("access_token"))
  spotifyApi.setAccessToken(VueCookies.get("access_token"));
}

export default{
  data() {
    return {
      watchVolume: 45,
      isPlay: true
    }
  },
  methods: {
    PausePlayback() {
      spotifyApi.pause()
      this.isPlay = !this.isPlay
    },
    PlayPlayback() {
      spotifyApi.play()
      this.isPlay = !this.isPlay
    },
    skipToPrevious() {
      spotifyApi.skipToPrevious()
    },
    skipToNext() {
      spotifyApi.skipToNext()
    }
  },
  watch: {
    watchVolume(volume) {
      spotifyApi.setVolume(volume)
    }
  }
}
</script>


<style>
@import url('https://fonts.googleapis.com/css2?family=Source+Sans+Pro:ital,wght@0,400;1,700&display=swap');
@import url('https://fonts.googleapis.com/icon?family=Material+Icons');
div {
  align-content: center;

  align-items: center;
}
.volume {
  padding-top: 50px;
  padding-right: 30px;
  padding-bottom: 50px;
  padding-left: 50px;
}
</style>