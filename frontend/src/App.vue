<template>
  <div class="h-full w-full flex flex-col">
    <Header class="h-16 flex-none col"></Header>
    <div class="flex flex-nowrap">
      <active-users class="w-60 flex-none"></active-users>
      <div class="flex flex-col" style="min-width: 30rem;">
        <router-view style="height: calc(100vh - 9rem)"></router-view>
        <player class="h-20">player</player>
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
