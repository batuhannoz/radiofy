<template>
    <div class="text-white flex flex-wrap">
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-purple-700"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-SpotifyGreen"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-blue-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-green-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-pink-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-purple-700"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-SpotifyGreen"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-blue-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-green-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-pink-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-purple-700"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-SpotifyGreen"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-blue-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-green-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-pink-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-purple-700"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-SpotifyGreen"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-blue-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-green-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-pink-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-purple-700"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-SpotifyGreen"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-blue-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-green-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-pink-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-purple-700"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-SpotifyGreen"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-blue-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-green-600"></div>
      <div class="flex-initial m-auto my-4 aspect-square h-40 w-40 bg-pink-600"></div>
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