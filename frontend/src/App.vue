<template>
  <div class="h-full w-full flex flex-col">
    <Header class="h-16 flex-none col"></Header>
    <div class="flex flex-nowrap flex-grow">
      <active-users class="w-60 flex-none"></active-users>
      <div class="flex flex-col flex-grow relative" style="min-width: 30rem;">
        <router-view class="flex-grow"></router-view>
        <player class="absolute inset-x-0 bottom-0 h-20">player</player>
      </div>
      <Chat class="w-60 flex-none">Chat</Chat>
    </div>
  </div>
</template>

<script>
import ActiveUsers from './components/ActiveUsers.vue'
import Player from './components/Player.vue'
import Chat from  './components/Chat.vue'
import Header from './components/Header.vue'

export default {
  components: {
    ActiveUsers,
    Player,
    Chat,
    Header
  }
}

</script>

<style>


</style>