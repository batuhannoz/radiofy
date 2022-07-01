<template>
  <div class="text-white grid grid-cols-4 overflow-y-auto scroller">
    <!--Create New Club-->
    <button @click="RedirectCreatePage" class="flex flex-col hover:bg-[#2a2a2a] duration-300 ease-in-out m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
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
    <button @click="RedirectClub(items.code)" v-for="items in clubList" :key="items" class="group duration-300 ease-in-out hover:bg-[#2a2a2a] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <div class="relative h-[165px] w-[165px]">
        <img :src="items.image" class="relative group-hover:shadow-4xl rounded-md"/>
        <div class="transition-all ease-in-out duration-[400ms] group-hover:shadow-4xl transform opacity-0 group-hover:opacity-100 group-hover:bottom-1.5 absolute  group-hover:flex group-hover:block right-1 bottom-0 bg-SpotifyGreen w-11 h-11 rounded-full flex items-center justify-center">
          <play-icon class="text-black"/>
        </div>
      </div>
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
import PlayIcon from "@/assets/icon/Play.vue";

export default {
  components:{
    PlayIcon
  },
  computed: {
    ...mapGetters("auth", ["getAccessToken", "getRadiofyToken"]),
    ...mapGetters("player", ["getDeviceID"])
  },
  data() {
    return{
      clubList: null
    }
  },
  mounted() {
    axios.get('http://localhost:3000/clubs',{
      headers: {
        "Authorization": this.getRadiofyToken
    }}).then((res) => {
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

