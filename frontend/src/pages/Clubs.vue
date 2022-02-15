<template>
  <div class="mid-container">
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
    <li>abc</li>
    <li>def</li>
    <li>ghi</li>
  </div>
</template>

<script>
import VueCookies from "vue-cookies";
import {setAccessToken} from "@/services/SpotifyPlayer";

export default {
  beforeCreate() {
    if (VueCookies.get("access_token")) {
      setAccessToken()
      console.log("logged in spotify")
    }else{
      this.$router.replace('/login')
    }
  }
}
</script>

<style>

.mid-container{
  flex-wrap: wrap;
  overflow-y: auto;
  height: 89%;
}

</style>