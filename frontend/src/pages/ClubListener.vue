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
    ...mapGetters("player", ["getDeviceID"])
  },
  methods: {
    ...mapActions("player",["playSong"])
  },
  mounted() {
    this.socket = new WebSocket("ws://localhost:3000/club/" + this.id)
    this.socket.onmessage = (msg) => {
      //create a JSON object
      let jsonObject = JSON.parse(msg.data);
      this.playSong({albumID: jsonObject.albumID, position: jsonObject.position - 1, deviceID: this.getDeviceID})
    }
  },
}
</script>