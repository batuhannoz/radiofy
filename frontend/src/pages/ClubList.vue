<template>
  <div class="text-white grid grid-cols-4 overflow-y-auto scroller">
    <!--Create New Club-->
    <button @click="RedirectCreatePage" class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="../assets/image/plus_icon.png" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-bold line-clamp-1 truncate text-center w-full">
          Create New Club
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Create your own club and share your music with people.
        </p>
      </div>
    </button>
    <!--Active Clubs-->
    <button @click="RedirectClub(items.code)" v-for="items in clubList" :key="items" class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img :src="items.image" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-bold line-clamp-1 text-center w-full">
          {{items.name}}
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          {{items.description}}
        </p>
      </div>
    </button>
  </div>
</template>

<script>
import axios from "axios";
import {mapGetters} from 'vuex';

export default {
  computed: {
    ...mapGetters("auth", ["getAccessToken"]),
    ...mapGetters("player", ["getDeviceID"])
  },
  data() {
    return{
      clubList: null
    }
  },
  mounted() {
    axios.get('http://localhost:3000/club/list').then((res) => {
      console.log(res)
      this.clubList = res.data
    })
  },
  methods: {
    RedirectCreatePage() {
      this.$router.push("/create")
    },
    RedirectClub(clubCode) {
      console.log(clubCode)
      this.$router.push("/club/"+ clubCode)
    }
  }
}
</script>
