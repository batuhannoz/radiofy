<template>
  <div class="h-full w-full flex flex-col">
    <Header class="h-[5.3rem] flex-none col"></Header>
    <div class="flex flex-nowrap">
      <active-users class="w-[16.5rem] flex-none"></active-users>
      <div class="flex flex-col" style="min-width: 50rem;">
        <router-view style="height: calc(100vh - 5.3rem - 92px)"></router-view>
        <player style="height: 92px" class="w-full">player</player>
      </div>
      <Chat class="w-80 flex-none">Chat</Chat>
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
