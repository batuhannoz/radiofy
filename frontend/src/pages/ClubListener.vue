<template>
  <div class="text-white">
    listening...
  </div>
</template>

<script>

import {mapActions, mapGetters} from "vuex";

export default {
  props:["id"],
  computed: {
    ...mapGetters("auth", ["getRadiofyToken"]),
    ...mapGetters("player", ["getDeviceID"])
  },
  methods: {
    ...mapActions("player",["playSong"])
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:3000/club/id/listen"+"?token="+this.getRadiofyToken)
    this.socket.onmessage = (msg) => {
      this.socket.send("pong")
      let jsonObject = JSON.parse(msg.data);
      this.playSong({albumID: jsonObject.albumID, position: jsonObject.position - 1, deviceID: this.getDeviceID})
    }
  },
}
</script>