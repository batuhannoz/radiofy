<template>
  <search-song/>
</template>

<script>
import searchSong from "@/components/search/SearchSong";
import {mapGetters} from "vuex";
import axios from "axios";
export default {
  props: ["id"],
  components: {
    searchSong
  },
  data() {
  },
  computed: {
    currentSong() {
      return this.getSongID
    },
    ...mapGetters("auth", ["getRadiofyToken"]),
    ...mapGetters("player", ["getSongID", "getSongName", "getArtistName", "getImage", "getProgressMS", "getAlbumID", "getPosition"])
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
      axios.post("http://localhost:3000/club/" + this.id + "/change", request, {
        headers: {
          "Authorization": this.getRadiofyToken
        }
      })
    }
  },
}
</script>