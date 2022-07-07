<template>
  <div class="text-white">
    listening...
  </div>
</template>

<script>

import {mapActions, mapGetters} from "vuex";

export default {
  props: ["id"],
  computed: {
    ...mapGetters("auth", ["getRadiofyToken"]),
    ...mapGetters("player", ["getDeviceID", "getSongName"])
  },
  methods: {
    ...mapActions("player", ["playSong"])
  },
  mounted() {
    this.socket = new WebSocket("ws://192.168.1.127:3000/club/" + this.id + "/song"+"?token=" + this.getRadiofyToken)
    this.socket.onmessage = (msg) => {
      let jsonObject = JSON.parse(msg.data);
      if (jsonObject.type === "ping") {
        let pong = {
          "type": "pong"
        }
        this.wsConn.send(pong)
      }
      this.playSong({albumID: jsonObject.albumID, position: jsonObject.position - 1, deviceID: this.getDeviceID})
    }
  }
}
</script>