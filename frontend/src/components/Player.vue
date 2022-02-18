<template>
  <div class="bg-SpotifyPlayer text-white flex justify-between ">
    <div class="flex-grow-0 " style="height: 92px; width: 134px;">
      <img style="height: 92px; width: 92px;" src="https://i.scdn.co/image/ab67616d00001e02e4c3db3e7ebfc22bc080f334">
    </div>
    <div class="flex-grow-0 flex flex-col w-[45%] justify-center" style="max-width: 720px;">
      <div class="flex gap-x-2 mb-2 justify-center flex-grow-0">
        <button @click="Shuffle" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center">
          <ShuffleIcon v-if="isShuffle" class="text-SpotifyGreen"/>
          <ShuffleIcon v-else class="text-[#b3b3b3]"/>
        </button>
        <button @click="SkipToPrev" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]" >
          <previous/>
        </button>
        <button @click="PlayPause" class="bg-white text-black hover:scale-[1.06] rounded-full w-8 h-8 flex-grow-0 flex justify-center items-center">
          <stop-icon v-if="isPlay"/>
          <play-icon v-else/>
        </button>
        <button @click="SkipToNext" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <next/>
        </button>
        <button @click="Repeat" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <repeat-icon v-if="isRepeat" class="text-SpotifyGreen"/>
          <repeat-icon v-else class="text-[#b3b3b3]" />
        </button>
      </div>
      <div class="w-full flex items-center mt-1.5 flex-grow-0 gap-x-2">
        <!--instant music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">00:00</div>
        <!--slider-->
        <div class="flex-grow">
          <vue-slider
              v-model="progress"
              :dot-size="15"
              :process-style="{ background: '#1db954' }"
              :bg-style="{ background: '#737575' }"
          ></vue-slider>
          </div>
        <!--total music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">00:00</div></div>
    </div>
    <div class="flex-grow-0 flex items-center mr-2 text-center" style="height: 92px; width: 134px;">
      <div class="flex-grow flex flex-nowrap items-center">
        <button @click="SetMute" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <mute-icon v-if="Volume===0"/>
          <VolumeLow v-else-if="Volume < 33"/>
          <VolumeMid v-else-if="Volume < 66"/>
          <VolumeHigh v-else/>
        </button>
        <vue-slider style="width: 102px;"
          v-model="Volume"
          :interval="5"
          :dot-size="15"
          :process-style="{ background: '#1db954' }"
          :bg-style="{ background: '#737575' }"
        ></vue-slider>
      </div>
    </div>
  </div>
</template>



<script>
import VueSlider from 'vue-slider-component'
import 'vue-slider-component/theme/antd.css'
import PlayIcon from './Icons/PlayIcon.vue'
import ShuffleIcon from './Icons/ShuffleIcon'
import StopIcon from './Icons/StopIcon.vue'
import Previous from "./Icons/PreviousIcon";
import Next from "./Icons/Next";
import RepeatIcon from './Icons/RepeatIcon.vue'
import VolumeHigh from './Icons/VolumeHigh.vue'
import MuteIcon from './Icons/MuteIcon.vue'
import VolumeLow from './Icons/VolumeLow.vue'
import VolumeMid from './Icons/VolumeMid.vue'

import {
  setVolume,
  repeat,
  shuffle,
  skipToPrevious,
  skipToNext,
  pause,
  play,
} from "@/services/SpotifyPlayer";
export default {
  components: {
    Next,
    Previous,
    PlayIcon,
    StopIcon,
    RepeatIcon,
    ShuffleIcon,
    VueSlider,
    VolumeHigh,
    MuteIcon,
    VolumeLow,
    VolumeMid
  },
  data() {
    return {
      progress: 0,
      isPlay: false,
      isShuffle: false,
      isRepeat: false,
      Volume: 50
    }
  },
  methods:{
    PlayPause(){
      if (this.isPlay){
        pause()
        this.isPlay = false
      }else {
        play()
        this.isPlay = true
      }
    },
    SkipToPrev() {
      skipToPrevious()
    },
    SkipToNext() {
      skipToNext()
    },
    Shuffle(){
      shuffle(!this.isShuffle)
      this.isShuffle = !this.isShuffle
    },
    Repeat(){
      if (this.isRepeat) {
        repeat('off')
      } else {
        repeat('track')
      }
      this.isRepeat = !this.isRepeat
    },
    SetMute() {
      this.Volume = 0
    }
  },
  watch: {
    Volume(vol) {
      setVolume(vol)
    }
  }
}


</script>