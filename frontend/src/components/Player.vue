<template>
  <div class="bg-SpotifyPlayer text-white flex justify-between ">
    <div class="flex-grow-0 " style="height: 92px; width: 134px;">
      <img v-if="GetImage" style="height: 92px; width: 92px;" v-bind:src="GetImage">
    </div>
    <div class="flex-grow-0 flex flex-col w-[45%] justify-center" style="max-width: 720px;">
      <div class="flex gap-x-2 mb-2 justify-center flex-grow-0">
        <button @click="Shuffle" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center">
          <ShuffleIcon v-if="GetIsShuffle" class="text-SpotifyGreen"/>
          <ShuffleIcon v-else class="text-[#b3b3b3]"/>
        </button>
        <button @click="SkipToPrev" class="hover:scale-[1.06] hover:text-white w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]" >
          <previous/>
        </button>
        <button @click="PlayPause" class="bg-white text-black hover:scale-[1.06] rounded-full w-8 h-8 flex-grow-0 flex justify-center items-center">
          <stop-icon v-if="GetIsPlay"/>
          <play-icon v-else/>
        </button>
        <button @click="SkipToNext" class="hover:scale-[1.06] hover:text-white w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <next/>
        </button>
        <button @click="Repeat" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <repeat-icon v-if="GetIsRepeat === 'off'" class="text-[#b3b3b3]"/>
          <repeat-icon v-else class="text-SpotifyGreen" />
        </button>
      </div>
      <div class="w-full flex items-center mt-1.5 flex-grow-0 gap-x-2">
        <!--instant music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">{{ MstoSecond(GetProgressMS) }}</div>
        <!--slider-->
        <div class="flex-grow">
          <vue-slider
              v-model=" Progress "
              @change="ChangeSpotifyProgress"
              :lazy="true"
              :max="GetDurationMS"
              :tooltip="'none'"
              :dot-size="15"
              :process-style="{ background: '#1db954' }"
              :bg-style="{ background: '#737575' }"
          ></vue-slider>
          </div>
        <!--total music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">{{ MstoSecond(GetDurationMS) }}</div></div>
    </div>
    <div class="flex-grow-0 flex items-center mr-2 text-center" style="height: 92px; width: 134px;">
      <div class="flex-grow flex flex-nowrap items-center">
        <button @click="SetMute" class="hover:scale-[1.06] hover:text-white w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
          <mute-icon v-if=" GetVolume === 0"/>
          <VolumeLow v-else-if=" GetVolume < 33"/>
          <VolumeMid v-else-if=" GetVolume < 66"/>
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

import {mapState, mapGetters, mapActions } from 'vuex'

import {
  skipToPrevious,
  skipToNext,
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
  computed: {
    ...mapState(["ProgressMS"]),
    ...mapGetters([
        "GetImage",
        "GetIsShuffle",
        "GetIsPlay",
        "GetIsRepeat",
        "GetProgressMS",
        "GetDurationMS",
        "GetVolume",
        "GetArtistName",
        "GetSongName",
    ]),
    change(){
      return this.GetProgressMS
    },
    Progress: {
      get() {
        return this.GetProgressMS
      },
      set(value) {
        this.ChangeProgressMsStore(value)
      },
    }
  },
  created() {
    setTimeout(() => {
      this.RefreshPlayer().then(() => {
        setTimeout(() => {
          if (this.GetIsPlay) {
            this.ProgressInterval = setInterval(() => {
              this.Progress += 1000
            }, 1000)
          }
        }, 1000)
      })
    }, 3000)
  },
  data() {
    return {
      ProgressInterval: null,
      Volume: 50,
    }
  },
  watch: {
    Volume(vol) {
      this.ChangeVolume(vol)
    },
    change(value) {
      if (value >= this.GetDurationMS) {
        if (this.GetIsPlay === true) {
          setTimeout(() => {
            this.RefreshPlayer()
          }, 3000)
        }
      }
    }
  },
  ...mapGetters([
    "GetImage",
    "GetIsShuffle",
    "GetIsPlay",
    "GetIsRepeat",
  ]),
  methods: {
    ...mapActions([
        "ChangeImage",
        "ChangeProgressMsStore",
        "ChangeProgressMsSpotify",
        "RefreshPlayer",
        "PlaySong",
        "PauseSong",
        "EnableShuffle",
        "DisableShuffle",
        "EnableRepeat",
        "DisableRepeat",
        "ChangeVolume"
    ]),
    SkipToPrev() {
      skipToPrevious()
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
      clearInterval(this.ProgressInterval)
      this.Progress = 0
      this.ProgressInterval = setInterval(() => {
        this.Progress += 1000
      }, 1000)
    },
    SkipToNext() {
      skipToNext()
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
      clearInterval(this.ProgressInterval)
      this.Progress = 0
      this.ProgressInterval = setInterval(() => {
        this.Progress += 1000
      }, 1000)
    },
    PlayPause() {
      if (this.GetIsPlay) {
        this.PauseSong()
        clearInterval(this.ProgressInterval)
      } else {
        this.PlaySong()
        this.ProgressInterval = setInterval(() => {
          this.Progress += 1000
        }, 1000)
      }
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
    },
    Shuffle() {
      if (this.GetIsShuffle) {
        this.DisableShuffle()
      }else {
        this.EnableShuffle()
      }
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
    },
    Repeat() {
      if (this.GetIsRepeat === 'off') {
        this.EnableRepeat()
      } else {
        this.DisableRepeat()
      }
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
    },
    SetMute() {
      this.Volume = 0
      this.ChangeVolume(0)
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
    },
    MstoSecond(ms) {
      return Math.floor(ms / 60000) + ":" + (((ms % 60000) / 1000).toFixed(0) < 10 ? '0' : '') + ((ms % 60000) / 1000).toFixed(0);
    },
    ChangeSpotifyProgress(value) {
      this.ChangeProgressMsSpotify(value)
      setTimeout(() => {
        this.RefreshPlayer()
      }, 3000)
    }
  }
}

</script>