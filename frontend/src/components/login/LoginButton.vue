<template>
  <div class="text-center w-full h-full">
    <button @click="RedirectLoginPage" class="font-medium px-4 py-2 hover:opacity-80 bg-SpotifyGreen hover:scale-105 rounded-full">
      Login With Spotify
    </button>
  </div>
</template>
<script>
import axios from "axios";

export default {
  methods: {
    RedirectLoginPage() {
    axios.get("http://192.168.1.127:3000/auth_url").then((res) => {
      window.location.href = res.data.url;
    })
    }
  }
};

</script>