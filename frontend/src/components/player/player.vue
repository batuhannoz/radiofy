<template>
  <div class="bg-SpotifyPlayer text-white flex justify-between ">
    <!--current music image-->
    <div class="flex-grow-0 " style="height: 92px; width: 140px;">
      <img v-if="getImage" style="height: 92px; width: 92px;" :src="getImage"/>
    </div>
    <!--player buttons-->
    <div class="flex-grow-0 flex flex-col w-[45%] justify-center" style="max-width: 720px;">
      <div class="flex gap-x-2 mb-2 justify-center flex-grow-0">
        <shuffle/>
        <previous/>
        <play v-if="getIsPlay"/>
        <Stop v-else/>
        <next/>
        <repeat/>
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
              :tooltip="'none'"
              :dot-size="15"
              :process-style="{ background: '#1db954' }"
              :bg-style="{ background: '#737575' }"
          ></vue-slider>
        </div>
        <!--total music duration-->
        <div class="text-[0.688rem] text-white text-opacity-70">{{MStoSecond(getDurationMS)}}</div>
      </div>
    </div>
    <!--volume control & buttons-->
    <div class="flex-grow-0 flex items-center mr-2 text-center" style="height: 92px; width: 140px;">
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
        ></vue-slider>
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
    return{
      Progress: 0,
      Volume: 0
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
      "getVolume"
    ]),
  },
  methods: {
    ...mapActions("player",[
      "playSong",
      "pauseSong",
      "setVolume",
      "setShuffle",
      "enableRepeat",
      "disableShuffle",
      "disableRepeat",
      "enableShuffle"]),
    MStoSecond(ms) {
      return Math.floor(ms / 60000) + ":" + (((ms % 60000) / 1000).toFixed(0) < 10 ? '0' : '') + ((ms % 60000) / 1000).toFixed(0);
    },
    play() {
      this.playSong()
    }
  }
}
</script>