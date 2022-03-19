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
      <div class="relative bg-[#1E1E1E] mt-1.5 mx-1.5 mb-4 rounded-xl" style="height: calc(100% - 32px)">
        <div class="m-2 flex flex-row">
          <div class="text-SpotifyGreen">
            Harlan Loflin:
          </div>
          <div>
             hi!
          </div>
        </div>
        <div class="m-2 flex flex-row">
          <div class="text-purple-700">
            Jonathan Weddel:
          </div>
          <div>
            hello
          </div>
        </div>
        <div class="m-2 flex flex-row">
          <div class="text-blue-700">
            Marcell Brodrick:
          </div>
          <div>
            hello
          </div>
        </div>
        <div class="m-2">
          <div class="text-pink-700">
            Erasmo Piccirillo:
          </div>
          Anyone want to suggest a song?
        </div>
        <div class="m-2">
          <div class="text-yellow-700">
            Marcell Brodrick:
          </div>
          I'm listening to NF these days, check it out if you want.
        </div>
        <div class="m-2">
          <div class="text-purple-700">
            Jonathan Weddel:
          </div>
          if you want i can open NF club so we can listen together.
        </div>
        <div class="flex flex-row inset-x-0 bottom-0 absolute">
          <input class="shadow appearance-none border rounded-l-xl w-full py-2 px-3 text-gray-700 leading-tight" type="text" placeholder="Write Message">
          <button v-bind:class="{[currentChatColor]: true}" class="inline-block px-3 py-2 text-white font-medium text-xs leading-tight uppercase rounded-r-xl shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 active:shadow-lg transition duration-150 ease-in-out flex items-center" type="button" id="button-addon2">
            <send-icon/>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import WorldIcon from "@/assets/icon/World.vue";
import FriendsIcon from "@/assets/icon/Friends.vue";
import ChatIcon from "@/assets/icon/Chat.vue";
import SendIcon from  "@/assets/icon/Send.vue";
export default {
  components: {
    WorldIcon,
    FriendsIcon,
    ChatIcon,
    SendIcon
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
      this.chatHeader = "Global Chat"
      this.currentChatColor = "bg-SpotifyGreen"
      this.currentChat = "global"
    },
    SetFriends() {
      this.chatHeader = "Friends"
      this.currentChatColor = "bg-purple-700"
      this.currentChat = "friend"
    },
    SetClub() {
      this.chatHeader = "Club Chat"
      this.currentChatColor = "bg-[#545454]"
      this.currentChat = "club"
    },
    SendMessage() {

    }
  }
}

</script>