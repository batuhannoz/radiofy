<template>
  <search-song/>
</template>

<script>
import searchSong from "@/components/search/SearchSong";
import {mapGetters} from "vuex";
export default {
  props: ["id"],
  components: {
    searchSong
  },
  mounted() {
    this.wsConn = new WebSocket("ws://localhost:3000/club/" + this.id + "/change"+"?token=" + this.getRadiofyToken)
    this.wsConn.onmessage = (msg) => {
      let jsonObject = JSON.parse(msg.data);
      if (jsonObject.type === "ping") {
        let pong = {
          "type": "pong"
        }
        this.wsConn.send(pong)
      }
    }
  },
  data() {
    return {
      wsConn: null
    }
  },
  computed: {
    ...mapGetters("auth", ["getRadiofyToken"]),
    ...mapGetters("player", ["getSongID", "getSongName", "getArtistName", "getImage", "getProgressMS", "getAlbumID", "getPosition"]),
    currentSong() {
      return this.getSongID
    },
  },
  watch: {
    currentSong(songID) {
      let request = {
        "albumID": this.getAlbumID,
        "position": this.getPosition,
        "clubID": this.id,
        "songID": songID,
        "songName": this.getSongName,
        "artistName": this.artistName,
        "image": this.getImage,
        "progressMS": this.getProgressMS,
      }
      this.wsConn.send(JSON.stringify(request))
    }
  },
}
</script>