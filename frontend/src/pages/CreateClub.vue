<template>
  <div class="text-white flex justify-center items-center">
    <div class="bg-[#252525] shadow-md rounded-lg px-8 pt-6 pb-8">
      <div class="mb-4">
        <label class="block text-[#b3b3b3] text-sm font-bold mb-2">
          Club Name
        </label>
        <input v-model="clubName" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight" type="text" placeholder="Club Name">
      </div>
      <div class="mb-4">
        <label class="block text-[#b3b3b3] text-sm font-bold mb-2">
          Club Description
        </label>
        <input v-model="clubDescription" class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight" type="text" placeholder="Club Description">
      </div>
      <div class="flex justify-center items-center mt-6">
        <button @click="createClub" class="tracking-wide font-medium px-4 py-2 hover:opacity-80 bg-SpotifyGreen hover:scale-[1.02] rounded-full">
          Create Club
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import axios from "axios";

export default {
  data() {
    return {
      clubName: "",
      clubDescription: "",
    };
  },
  computed: {
    ...mapGetters("player", ["getImage"])
  },
  methods: {
    createClub() {
        let clubReq = {
        "image": this.getImage,
        "name": this.clubName,
        "description": this.clubDescription
      }
      axios.post("http://localhost:3000/create/club", clubReq).then((res) => {
        this.$router.push("/club/leader/" + res.data.code)
      })
    }
  }
}
</script>


