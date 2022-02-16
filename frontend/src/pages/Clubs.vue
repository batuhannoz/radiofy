<template>
    <div class="text-white">
      Clubs
    </div>
</template>

<script>
import VueCookies from "vue-cookies";
import {setAccessToken} from "@/services/SpotifyPlayer";

export default {
  beforeCreate() {
    if (VueCookies.get("access_token")) {
      setAccessToken()
    }else{
      this.$router.replace('/login')
    }
  }
}
</script>

<style>

</style>