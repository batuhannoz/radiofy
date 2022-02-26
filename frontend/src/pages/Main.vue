<template>
  <div class="text-white grid grid-cols-4 overflow-y-auto scroller ">
    <!--Create New Club-->
    <button @click="RedirectCreatePage" class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="../assets/image/plus_icon.png" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 truncate text-center w-full">
          Create New Club
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Create your own club and share your music with people.
        </p>
      </div>
    </button>
    <!--Active Clubs-->
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/993cdbae6bf9c21aa653ada2d153c8f05f1b842e" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 text-center w-full">
          Science Club Lorem ipsum dolor sit amet, consectetur
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          listening to science podcasts Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        </p>
      </div>
    </button>
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/ab67616d00001e025675e83f707f1d7271e5cf8a" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 text-center w-full">
          Imagine Dragons Lorem ipsum dolor sit amet, consectetur
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          We love Imagine Dragons Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        </p>
      </div>
    </button>
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/ab67616d00001e026ca5c90113b30c3c43ffb8f4" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 text-center w-full">
          Eminem Club Lorem ipsum dolor sit amet, consectetur
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Eminem and Dr. Dre Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        </p>
      </div>
    </button>
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/ab67706f0000000250f34f0e7a0b73effe02fa6f" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 text-center w-full">
          Broken Hearts Club ipsum dolor sit amet, consectetur
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Broken Hearts Club Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        </p>
      </div>
    </button>
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/ab67706c0000bebb443715c10bcde6c10f733874" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 text-center w-full">
          Lorem ipsum dolor sit amet, consectetur
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
        </p>
      </div>
    </button>
    <button class="hover:opacity-70 hover:scale-[1.02] flex flex-col m-auto my-2 h-64 w-48 bg-SpotifyPlayer rounded-lg p-3.5">
      <img src="https://i.scdn.co/image/ab67706f000000027dc2a0fb4bfe0d07294f685e" class="rounded-md"/>
      <div class="h-full w-full flex justify-center mt-1 flex-col items-center">
        <div class="font-mono line-clamp-1 text-center w-full">
          Lorem ipsum dolor sit amet, consectetur
        </div>
        <p class="text-center line-clamp-2 w-full h-9 text-xs text-[#b3b3b3]">
          Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor
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
  },
  methods: {
    RedirectCreatePage() {
      this.$router.push("/create")
    }
  }
}
</script>
