<template>
  <div class="bg-SpotifyPlayer flex text-white">
    <div class="flex justify-around flex-none flex-row h-10 w-full mt-2">
      <button @click="SetGlobal" v-if="currentChat === 'global'" class="flex bg-SpotifyGreen items-center justify-center rounded-t-full h-12 w-[52px] hover:scale-105  focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 transition duration-150 ease-in-out">
        <world-icon/>
      </button>
      <button @click="SetGlobal" v-else class="flex bg-SpotifyGreen items-center justify-center rounded-full h-10 w-10 hover:scale-110 shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 transition duration-150 ease-in-out">
        <world-icon/>
      </button>
      <button @click="SetFriends" v-if="currentChat === 'friend'" class="flex bg-purple-700 justify-center items-center rounded-t-full h-12 w-[52px] hover:scale-105  focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 transition duration-150 ease-in-out">
        <friends-icon/>
      </button>
      <button @click="SetFriends" v-else class="flex bg-purple-700 justify-center items-center rounded-full h-10 w-10 hover:scale-110 shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 transition duration-150 ease-in-out">
        <friends-icon/>
      </button>
      <button @click="SetClub" v-if="currentChat === 'club'" class="flex  justify-center items-center bg-[#545454] rounded-t-full h-12 w-[52px] hover:scale-105  focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 transition duration-150 ease-in-out">
        <chat-icon/>
      </button>
      <button @click="SetClub" v-else class="flex justify-center items-center bg-[#545454] rounded-full h-10 w-10 hover:scale-110 shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 transition duration-150 ease-in-out">
        <chat-icon/>
      </button>
    </div>
    <div class="flex flex-col w-full" style="height: calc(100% - 48px)">
      <div v-bind:class="{[currentChatColor]: true}" class="w-full h-10 mt-2 font-semibold pt-1.5 text-center rounded-t-xl">{{chatHeader}}</div>
      <div class="relative bg-[#1E1E1E] mt-1.5 mx-1.5 mb-4 rounded-xl" style="height: calc(100% - 4rem)">
        <div class="h-full">
          <!--chat type check-->
          <club-chat v-if="currentChat === 'club'"/>
          <friends-chat v-else-if="currentChat === 'friend'"/>
          <global-chat v-else/>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import WorldIcon from "@/assets/icon/World.vue";
import FriendsIcon from "@/assets/icon/Friends.vue";
import ChatIcon from "@/assets/icon/Chat.vue";
import ClubChat from "../chat/Club.vue"
import GlobalChat from "../chat/Global.vue"
import FriendsChat from "../chat/Friends.vue"
export default {
  components: {
    WorldIcon,
    FriendsIcon,
    ChatIcon,
    ClubChat,
    GlobalChat,
    FriendsChat
  },
  data() {
    return{
      chatHeader: "Global Chat",
      currentChat: "global",
      currentChatColor: "bg-SpotifyGreen"
    }
  },
  methods: {
    SetGlobal() {
      console.log(this.currentChat)
      this.chatHeader = "Global Chat"
      this.currentChatColor = "bg-SpotifyGreen"
      this.currentChat = "global"
    },
    SetFriends() {
      console.log(this.currentChat)
      this.chatHeader = "Friends"
      this.currentChatColor = "bg-purple-700"
      this.currentChat = "friend"
    },
    SetClub() {
      console.log(this.currentChat)
      this.chatHeader = "Club Chat"
      this.currentChatColor = "bg-[#545454]"
      this.currentChat = "club"
    },
  }
}

</script>