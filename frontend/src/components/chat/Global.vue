<template>
  <div class="h-full">
    <div class="overflow-y-auto" style="height: calc(100% - 40px)">
      <div v-for="message in ReceviedMessages" :key="message.message" class="w-72 h-20">

        <div v-if="message.displayName === this.getDisplayName" class="ml-2 mt-1 w-72 h-20 flex flex-row-reverse">
          <img class="w-10 h-10 rounded-full" :src="message.image">
          <chat-triangle-right class="mr-1 mt-3 text-[#4f4f4f]"/>
          <div class="px-1.5 ml-2 py-0.5 mt-3 w-60 h-14 bg-[#4f4f4f] rounded-b-lg rounded-l-lg">
            {{message.message}}
          </div>
        </div>

        <div v-else class="ml-2 mt-1 w-72 h-20 flex flex-row">
          <img class="w-10 h-10 rounded-full" :src="message.image">
          <chat-triangle-left class="ml-1 mt-3 text-[#4f4f4f]"/>
          <div class="px-1.5 mr-2 py-0.5 mt-3 h-14 w-60 bg-[#4f4f4f] rounded-b-lg rounded-r-lg">
            {{message.message}}
          </div>
        </div>

      </div>
    </div>
    <div class="flex flex-row">
      <input v-on:keyup.enter="SendMessage" v-model="Message" class="shadow appearance-none border rounded-l-xl w-full py-2 px-3 text-gray-700 leading-tight" type="text" placeholder="Write Message">
      <button @click="SendMessage" class="bg-SpotifyGreen inline-block px-3 py-2 text-white font-medium text-xs leading-tight uppercase rounded-r-xl shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 active:shadow-lg transition duration-150 ease-in-out flex items-center" type="button" id="button-addon2">
        <send-icon/>
      </button>
    </div>
  </div>
</template>

<script>
import SendIcon from  "@/assets/icon/Send.vue";
import ChatTriangleLeft from "@/assets/icon/ChatTriangleLeft";
import ChatTriangleRight from "@/assets/icon/ChatTriangleRight";
import {mapGetters} from "vuex";


export default {
  components:{
    SendIcon,
    ChatTriangleLeft,
    ChatTriangleRight
  },
  computed: {
    ...mapGetters("user", ["getDisplayName"]),
    ...mapGetters("auth", ["getRadiofyToken"])
  },
  data(){
    return{
      GlobalChatSocket: null,
      Message: "",
      ReceviedMessages: []
    }
  },
  mounted() {
    this.GlobalChatSocket = new WebSocket("ws://192.168.1.127:3000/chat/global" + "?token=" + this.getRadiofyToken)
    this.GlobalChatSocket.onmessage = (msg) => {
      let jsonObject = JSON.parse(msg.data);
      if (jsonObject.type === "ping") {
        let pong = {
          "type": "pong"
        }
        this.wsConn.send(pong)
      }
      this.ReceviedMessages.push(jsonObject)
    }
  },
  methods: {
    SendMessage() {
      let msg = {
        message: this.Message
      }
      this.GlobalChatSocket.send(JSON.stringify(msg))
      this.Message = ""
    }
  }
}
</script>