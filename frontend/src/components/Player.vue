<template>
  <div class="bg-SpotifyPlayer text-white flex justify-between">
    <div class="flex-grow-0" style="height: 92px; width: 92px;">
      <img style="height: 92px; width: 92px;" src="https://i.scdn.co/image/ab67616d00001e02e4c3db3e7ebfc22bc080f334">
    </div>
    <div class="flex-grow-0 flex flex-col w-[45%] justify-center" style="max-width: 720px;">
      <div class="flex gap-x-2 mb-2 justify-center flex-grow-0">
        <button class="hover:scale-[1.06] rounded-full w-8 h-8 flex-grow-0 flex justify-center items-center">
          <stop-icon/>
        </button>
        <button class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]" >
          <previous/>
        </button>
        <button class="bg-white text-white hover:scale-[1.06] rounded-full w-8 h-8 flex-grow-0 flex justify-center items-center">
          <stop-icon/>
        </button>
        <button class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <next/>
        </button>
        <button class="hover:scale-[1.06] rounded-full w-8 h-8 flex-grow-0 flex justify-center items-center">
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
import Previous from "@/components/Icons/Previous";
import Next from "@/components/Icons/Next";

export default {
  components: {
    Next,
    Previous,
    //PlayIcon,
    StopIcon
  }
}


</script>