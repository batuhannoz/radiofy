<template>
  <div class="text-white grid grid-cols-4 overflow-y-auto scroller ">
    <!--Create New Club-->
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="../assets/image/plus_icon.png" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 truncate text-center w-full">
          Create New Club
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Create your own club and share your music with people.
        </p>
      </div>
      <!--Active Clubs-->
    </button>
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/ab67706c0000da84443715c10bcde6c10f733874" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 truncate text-center w-full">
          Lana del rey club
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Lana del rey seviyorsan gel brom qwtrqewrqewrqewtqetq qewrqwe qew qwer qwer qwer qwer qwer qwer
        </p>
      </div>
    </button>
  </div>
</template>

<script>
import {mapGetters} from 'vuex';

export default {
  computed: {
    ...mapGetters("auth", ["getAccessToken"]),
    ...mapGetters("player", ["getDeviceID"])
  }
}
</script>
