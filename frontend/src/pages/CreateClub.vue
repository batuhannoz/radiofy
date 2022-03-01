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
export default {
  data() {
    return {
      receivedDes: "",
      receivedName: "",
      clubName: "",
      clubDescription: "",
      socket: null
    };
  },/*
  mounted() {
    this.socket = new WebSocket("ws://localhost:3000/ws")
    this.socket.onmessage = (msg) => {
      //create a JSON object
      let jsonObject = JSON.parse(msg.data);
      this.receivedName = jsonObject.ClubName;
      this.receivedDes = jsonObject.ClubDescription;
      console.log("Received data string");
    }
  },
  methods: {
    createClub() {
      console.log(this.clubName);
      console.log(this.clubDescription);
      let msg = {
        "ClubName": this.clubName,
        "ClubDescription": this.clubDescription
      }
      this.socket.send(JSON.stringify(msg))
      this.$router.push("/search")
    }
  }*/
  methods: {
    createClub() {

    }
  }
}
</script>


