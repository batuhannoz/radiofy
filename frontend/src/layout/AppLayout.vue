<template>
  <div class="flex-col flex h-full w-full">
    <app-header class="h-[5.3rem] flex-none col"/>
    <div class="flex flex-nowrap">
      <active-users class="w-[16.5rem] flex-none hidden xl:block"/>
      <div class="flex flex-col flex-grow" style="min-width: 40rem">
        <router-view style="height: calc(100vh - 5.3rem - 92px)"/>
        <player style="height: 92px" class="w-full"/>
      </div>
      <chat class="w-80 flex-none hidden md:block"/>
    </div>
  </div>
</template>

<script>
import appHeader from '../components/app/Header.vue';
import chat from '../components/app/Chat.vue';
import activeUsers from '../components/app/ActiveUsers.vue';
import player from "../components/player/player.vue";

export default {
  components: {
    appHeader,
    chat,
    activeUsers,
    player
  }
};
</script>