<template>
  <div class="bg-SpotifyPlayer text-white flex justify-between">
    <div class="flex-grow-0" style="height: 92px; width: 92px;">
      <img style="height: 92px; width: 92px;" src="https://i.scdn.co/image/ab67616d00001e02e4c3db3e7ebfc22bc080f334">
    </div>
    <div class="flex-grow-0 flex flex-col w-[45%] justify-center" style="max-width: 720px;">
      <div class="flex gap-1 mb-2 justify-center flex-grow-0">
        <button class="bg-white hover:scale-[1.06] rounded-full w-8 h-8 flex-grow-0 flex justify-center items-center text-black">
          <stop-icon/>
        </button>
      </div>
      <div class="flex-grow-0">
        player time
      </div>
    </div>
    <div class="flex-grow-0 flex items-center text-center">
      volume
    </div>
  </div>
</template>

<script>
//import PlayIcon from './Icons/PlayIcon.vue'
import StopIcon from './Icons/StopIcon.vue'

export default {
  components: {
    //PlayIcon,
    StopIcon
  }
}


</script>