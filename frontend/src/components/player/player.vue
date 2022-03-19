<template>
  <div class="bg-SpotifyPlayer text-white flex justify-between ">
    <!--current music image-->
    <div class="flex-grow-0 flex flex-nowrap" style="height: 92px; width: 160px;">
      <img v-if="getImage" style="height: 92px; width: 92px;" :src="getImage"/>
    </div>
    <!--player buttons-->
    <div class="flex-grow-0 flex flex-col w-[45%] justify-center" style="max-width: 720px;">
      <div class="flex gap-x-2 mb-2 justify-center flex-grow-0">
        <shuffle @click="disableShuffle" v-if="getIsShuffle" class="bg-SpotifyGreen"/>
        <shuffle @click="enableShuffle" v-else/>
        <previous @click="skipToPrevious"/>
        <Stop @click="pauseSong" v-if="getIsPlay"/>
        <play @click="play" v-else/>
        <next @click="skipToNext"/>
        <repeat @click="disableRepeat" v-if="getIsRepeat === 'track'" class="bg-SpotifyGreen"/>
        <repeat @click="enableRepeat" v-else/>
      </div>
      <div class="w-full flex items-center mt-1.5 flex-grow-0 gap-x-2">
        <!--instant music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">{{MStoSecond(getProgressMS)}}</div>
        <!--slider-->
        <div class="flex-grow">
          <vue-slider
              v-model="Progress"
              :lazy="true"
              :max="getDurationMS"
              @error="error"
              :tooltip="'none'"
              :dot-size="15"
              :process-style="{ background: '#1db954' }"
              :bg-style="{ background: '#737575' }"
          />
        </div>
        <!--total music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">{{MStoSecond(getDurationMS)}}</div>
      </div>
    </div>
    <!--volume control & buttons-->
    <div class="flex-grow-0 flex items-center mr-2 text-center" style="height: 92px; width: 160px;">
      <div class="flex-grow flex flex-nowrap items-center justify-between">
        <mute v-if="getVolume === 0"/>
        <volume-low v-else-if="getVolume <= 33"/>
        <volume-mid v-else-if="getVolume <= 66"/>
        <volume-high v-else/>
        <vue-slider style="width: 102px;"
                    v-model="Volume"
                    :interval="5"
                    :dot-size="15"
                    :process-style="{background: '#1db954'}"
                    :bg-style="{background: '#737575'}"
        />
      </div>
    </div>
  </div>
</template>

<script>
import VueSlider from 'vue-slider-component'
import 'vue-slider-component/theme/antd.css'
import Repeat from "@/assets/button/Repeat";
import Mute from "@/assets/button/Mute";
import VolumeLow from "@/assets/button/VolumeLow";
import VolumeMid from "@/assets/button/VolumeMid";
import VolumeHigh from "@/assets/button/VolumeHigh";
import Next from "@/assets/button/Next";
import Stop from "@/assets/button/Stop";
import Play from "@/assets/button/Play";
import Previous from "@/assets/button/Previous";
import Shuffle from '@/assets/button/Shuffle.vue';
import {mapActions, mapGetters} from "vuex";

export default {
  components: {
    Shuffle,
    Previous,
    Mute,
    VolumeLow,
    VolumeHigh,
    VolumeMid,
    Repeat,
    Next,
    Stop,
    Play,
    VueSlider
  },
  data() {
    return {
      ProgressInterval: null,
    }
  },
  watch:{
    isPlay(isPlay) {
      if(isPlay) {
        clearInterval(this.ProgressInterval)
        this.ProgressInterval = setInterval(() => {
          this.setProgress(this.Progress + 1000)
        }, 1000)
      }else{
        clearInterval(this.ProgressInterval)
      }
    }
  },
  computed: {
    ...mapGetters("player", [
      "getImage",
      "getArtistName",
      "getSongName",
      "getDurationMS",
      "getIsRepeat",
      "getProgressMS",
      "getIsShuffle",
      "getIsPlay",
      "getVolume",
    ]),
    isPlay: {
      get() {
        return this.getIsPlay
      }
    },
    Progress: {
      get() {
        return this.getProgressMS
      },
      set(value) {
        this.seekThePosition(value)
      },
    },
    Volume: {
      get() {
        return this.getVolume
      },
      set(volume) {
        this.setVolume(volume)
      }
    }
  },
  methods: {
    ...mapActions("player",[
      "play",
      "pauseSong",
      "setVolume",
      "setShuffle",
      "enableRepeat",
      "disableShuffle",
      "disableRepeat",
      "enableShuffle",
      "seekThePosition",
      "skipToNext",
      "skipToPrevious",
      "setProgress",
      "refreshPlayer"]),
    MStoSecond(ms) {
      return Math.floor(ms / 60000) + ":" + (((ms % 60000) / 1000).toFixed(0) < 10 ? '0' : '') + ((ms % 60000) / 1000).toFixed(0);
    },
    error(errorType) {
      /*
      * const ERROR_TYPE = {
      *   VALUE: 1,
      *   INTERVAL: 2,
      *   MIN: 3,
      *   MAX: 4,
      *   ORDER: 5,
      * }
      * */
      if(errorType === 4) {
        setTimeout(() => {
          this.refreshPlayer()
        },1500)
      }
    }
  }
}
</script>