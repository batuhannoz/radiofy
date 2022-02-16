<template>
  <div class="justify-center text-center">
    <button class="px-4 py-2 bg-SpotifyGreen hover:scale-105 rounded-full" @click="loginSpotify">Login With Spotify</button>
  </div>
</template>

<script>
export default {
  methods: {
    // redirect spotify login page
    loginSpotify() {
      const url = "https://accounts.spotify.com/authorize";
      let params = {
        response_type: "code",
        client_id: "d97ea09987ab46c2bb0b4fa4eaae55e1",
        scope:
            "user-read-private user-read-email playlist-modify-private playlist-read-collaborative playlist-read-private playlist-modify-public user-read-recently-played user-top-read user-read-playback-position user-read-playback-state user-modify-playback-state user-modify-playback-state",
        redirect_uri: "http://localhost:8080/token",
        show_dialog: true,
      };
      params = new URLSearchParams(params).toString();
      window.location.href = `${url}?${params}`;
    }
  }
}
</script>

<style>

</style>