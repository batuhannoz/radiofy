<template>
  <div class="flex flex-col bg-SpotifyBlack text-white overflow-y-auto scroller">
    <div v-for="listener in listeners" :key="listener.displayName" class="bg-SpotifyPlayer h-12 w-auto my-2 mx-[3px] rounded-full flex flex-row group">
      <img :src="listener.image" class="rounded-full m-0.5" >
      <div class="ml-2 flex-col flex py-0.5 pl-1 pr-3">
        <div class="font-bold text-[#b3b3b3]" style="max-width: 145px">{{listener.displayName}}</div>
        <div class="line-clamp-1 text-xs text-[#b3b3b3]" style="max-width: 145px">{{listener.songName}} - {{listener.artistName}}</div>
      </div>
        <div v-if="listener.albumID !== ''" @click="playListenerSong(listener.albumID, listener.position)" class="opacity-0 flex rounded-full group-hover:opacity-100 transition-all bg-SpotifyGreen duration-300 my-2 mr-2 h-8 w-8 items-center justify-center">
        <play-icon/>
      </div>
    </div>
  </div>
</template>

<script>
import PlayIcon from "@/assets/icon/Play";
import {mapActions ,mapGetters} from "vuex";

export default {
  components:{
    PlayIcon
  },
  computed: {
    ...mapGetters("auth", ["getRadiofyToken"]),
    ...mapGetters("player", ["getSongID", "getSongName", "getArtistName", "getDeviceID", "getPosition", "getAlbumID"]),
    currentSong() {
      return this.getSongID
    }
  },
  mounted() {
    this.wsConn = new WebSocket("ws://192.168.1.127:3000/listeners" + "?token=" + this.getRadiofyToken)
    this.wsConn.onmessage = (listeners) => {
      let jsonObject = JSON.parse(listeners.data);
      if (jsonObject.type === "ping") {
        let pong = {
          "type": "pong"
        }
        this.wsConn.send(pong)
      }
      this.listeners = jsonObject
    }
  },
  data() {
    return {
      wsConn: null,
      listeners: null
    }
  },
  watch: {
    currentSong(songID) {
      let request = {
        "albumID": this.getAlbumID,
        "position": this.getPosition,
        "songID": songID,
        "songName": this.getSongName,
        "artistName": this.getArtistName,
      }
      this.wsConn.send(JSON.stringify(request))
    }
  },
  methods: {
    ...mapActions("player", ["playSong"]),
    playListenerSong(albumID, position) {
      this.playSong({albumID: albumID, position: position-1, deviceID: this.getDeviceID})
    }
  }
}
</script>