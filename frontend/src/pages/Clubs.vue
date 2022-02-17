<template>
    <div class="text-white flex flex-wrap overflow-y-auto scroller">
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
      </div>
      <div class="flex-initial flex flex-col m-auto my-6 aspect-square h-56 w-40 bg-SpotifyPlayer rounded-lg p-4">
        <div>
          <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
        </div>
        <p class="text-center justify-center">
          werwereqw
          qweqwe
        </p>
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

