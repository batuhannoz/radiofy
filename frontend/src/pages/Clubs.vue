<template>
    <div class="text-white grid grid-cols-4 overflow-y-auto scroller ">
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
      </div>
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da8426c06a0d8c06cb366c2c1be8" class="rounded-md"/>
        </div>
      </div>
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67616d0000b2733e0d05346299e86c58f75123" class="rounded-md"/>
        </div>
      </div>
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67616d00001e0211ef0d6febc70a5443eefab4" class="rounded-md"/>
        </div>
      </div>
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67616d00001e02d274d0df7dbc820730c84878" class="rounded-md"/>
        </div>
      </div>
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67616d00001e02fc8285b4f20f2636a27464bd" class="rounded-md"/>
        </div>
      </div>
      <div class=" flex flex-col m-auto my-2 h-60 w-48 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67616d00001e02d9b35d1c4d15c9de88b754a7" class="rounded-md"/>
        </div>
      </div>
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

