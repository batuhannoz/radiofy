<template>
  <div class="h-full w-full flex flex-col ">
    <!--navbar-->
    <div class="flex justify-center h-10">
      <div class="mb-3 xl:w-96">
        <div class="input-group relative flex flex-nowrap items-stretch w-full mb-4">
          <input v-model="searchName" type="search" class="form-control relative flex-auto min-w-0 block w-full px-2.5 py-1 text-base font-normal text-gray-700 bg-white bg-clip-padding border border-solid border-gray-300 rounded-l-full transition ease-in-out m-0 focus:text-gray-700 focus:bg-white focus:border-SpotifyGreen focus:border-2 focus:outline-none" placeholder="Search" aria-label="Search" aria-describedby="button-addon2">
          <button @click="searchTrack" class="btn inline-block px-5 py-2 bg-SpotifyGreen text-white font-medium text-xs leading-tight uppercase rounded-r-full shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 active:shadow-lg transition duration-150 ease-in-out flex items-center" type="button" id="button-addon2">
            <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" class="w-4" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
              <path fill="currentColor" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path>
            </svg>
          </button>
        </div>
      </div>
    </div>
    <!--item list-->
    <ul class="text-white grid overflow-y-auto scroller grid-cols-4">
      <button @click="playSearchedSong(items.album.uri, items.track_number)" v-for="items in searchedSongs" :key="items" class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
        <img :src="items.album.images[0].url" class="rounded-md"/>
        <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
          <div class="font-bold line-clamp-1 text-center w-full">
            {{items.name}}
          </div>
          <p class="text-center w-full h-9 text-xs text-[#b3b3b3]">
            {{items.album.artists[0].name}}
          </p>
        </div>
      </button>
    </ul>
  </div>
</template>

<script>
import {mapActions, mapGetters} from "vuex";
import {searchItem} from '@/api/spotify/player.js';
export default {
  data() {
    return {
      searchName: "",
      searchedSongs: null
    }
  },
  computed: {
    ...mapGetters("player",["getDeviceID"]),
  },
  methods: {
    ...mapActions("player", ["playSong"]),
    searchTrack() {
      searchItem(this.searchName, 48).then((res) => {
        this.searchedSongs = res.tracks.items
      })
    },
    playSearchedSong(albumID, position) {
      this.playSong({albumID: albumID, position: position-1, deviceID: this.getDeviceID})
    }
  }
}
</script>