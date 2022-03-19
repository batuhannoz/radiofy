<template>
  <div class="bg-purple-700 h-full overflow-y-auto">
    {{ReceviedMessage}}
    <div class="flex flex-row inset-x-0 bottom-0 absolute">
      <input v-model="Message" class="shadow appearance-none border rounded-l-xl w-full py-2 px-3 text-gray-700 leading-tight" type="text" placeholder="Write Message">
      <button @click="SendMessage" class="bg-purple-700 inline-block px-3 py-2 text-white font-medium text-xs leading-tight uppercase rounded-r-xl shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 active:shadow-lg transition duration-150 ease-in-out flex items-center" type="button" id="button-addon2">
        <send-icon/>
      </button>
    </div>
  </div>
</template>

<script>
import SendIcon from  "@/assets/icon/Send.vue";
export default {
  components:{
    SendIcon
  },
  data(){
    return{
      Message: "",
      ReceviedMessage: ""
    }
  },
  methods: {
    SendMessage() {
      this.ReceviedMessage += " " + this.Message
      this.Message = ""
    }
  }
}
</script>