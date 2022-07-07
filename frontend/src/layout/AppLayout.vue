<template>
  <div class="flex-col flex h-full w-full">
    <app-header class="h-[5.3rem] flex-none col"/>
    <div class="flex flex-nowrap">
      <active-users v-if="this.getRadiofyToken !== ''" class="w-[16.5rem] flex-none hidden xl:block" style="height: calc(100vh - 5.3rem)"/>
      <div class="flex flex-col flex-grow" style="min-width: 50rem">
        <router-view style="height: calc(100vh - 5.3rem - 92px)"/>
        <player style="height: 92px" class="w-full"/>
      </div>
      <chat v-if="this.getRadiofyToken !== ''" class="w-80 flex-none hidden md:block" style="height: calc(100vh - 5.3rem)"/>
    </div>
  </div>
</template>

<script>
import appHeader from '../components/app/Header.vue';
import chat from '../components/app/Chat.vue';
import activeUsers from '../components/app/ActiveUsers.vue';
import player from "../components/player/player.vue";
import {mapGetters} from "vuex";

export default {
  components: {
    appHeader,
    chat,
    activeUsers,
    player
  },
  computed: {
    ...mapGetters("auth", ["getRadiofyToken"])
  },
};
</script>