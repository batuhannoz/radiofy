<template>
  <div class="text-white">
    listening...
  </div>
</template>

<script>

import {mapActions, mapGetters} from "vuex";
import axios from "axios";

export default {
  props:["id"],
  computed: {
    ...mapGetters("auth", ["getRadiofyToken"]),
    ...mapGetters("player", ["getDeviceID", "getSongName"])
  },
  methods: {
    ...mapActions("player",["playSong"])
  },
  data() {
    return {
      interval: null
    }
  },
  mounted() {
    this.interval = setInterval(() => {
      axios.get("http://localhost:3000/club/"  + this.id + "/song", {
        headers: {
          "Authorization": this.getRadiofyToken
        }
      }).then((res) => {
        if (res.data.songName !== this.getSongName) {
          this.playSong({albumID: res.data.albumID, position: res.data.position-1, deviceID: this.getDeviceID})
        }
      })
    }, 15000)
  },
  beforeUnmount() {
    clearInterval(this.interval)
  }
}
</script>