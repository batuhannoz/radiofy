<template>
  <div class="bg-SpotifyPlayer flex text-white">
    <div class="flex justify-around flex-none flex-row h-10 w-full mt-2">
      <button class="flex bg-SpotifyGreen items-center justify-center rounded-t-full h-10 w-10">
        <svg style="width:24px;height:24px" viewBox="0 0 24 24">
          <path fill="currentColor" d="M17.9,17.39C17.64,16.59 16.89,16 16,16H15V13A1,1 0 0,0 14,12H8V10H10A1,1 0 0,0 11,9V7H13A2,2 0 0,0 15,5V4.59C17.93,5.77 20,8.64 20,12C20,14.08 19.2,15.97 17.9,17.39M11,19.93C7.05,19.44 4,16.08 4,12C4,11.38 4.08,10.78 4.21,10.21L9,15V16A2,2 0 0,0 11,18M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z" />
        </svg>
      </button >
      <button class="flex bg-purple-700 justify-center items-center rounded-full h-10 w-10">
        <svg style="width:24px;height:24px" viewBox="0 0 24 24">
          <path fill="currentColor" d="M12 4C14.2091 4 16 5.79086 16 8C16 10.2091 14.2091 12 12 12C9.79086 12 8 10.2091 8 8C8 5.79086 9.79086 4 12 4ZM12 14C16.4183 14 20 15.7909 20 18L20 20L4 20L4 18C4 15.7909 7.58172 14 12 14Z"></path>
        </svg>
      </button>
      <button class="flex justify-center items-center bg-[#A3A3A3] rounded-full h-10 w-10">
        <svg style="width:24px;height:24px" viewBox="0 0 24 24">
          <path fill="currentColor" d="M15,4V11H5.17L4,12.17V4H15M16,2H3A1,1 0 0,0 2,3V17L6,13H16A1,1 0 0,0 17,12V3A1,1 0 0,0 16,2M21,6H19V15H6V17A1,1 0 0,0 7,18H18L22,22V7A1,1 0 0,0 21,6Z" />
        </svg>
      </button>
    </div>
    <div class="flex flex-col w-full" style="height: calc(100% - 48px)">
      <div class=" w-full h-8 font-semibold pt-1 text-center rounded-t-xl bg-SpotifyGreen">Global Chat</div>
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
          <button class="btn inline-block px-3 py-2 bg-SpotifyGreen text-white font-medium text-xs leading-tight uppercase rounded-r-xl shadow-md hover:opacity-75 hover:shadow-lg focus:shadow-lg focus:outline-none focus:ring-0 active:opacity-50 active:shadow-lg transition duration-150 ease-in-out flex items-center" type="button" id="button-addon2">
            <svg viewBox="0 0 24 24" width="24" height="24" class=""><path fill="currentColor" d="M1.101 21.757 23.8 12.028 1.101 2.3l.011 7.912 13.623 1.816-13.623 1.817-.011 7.912z"></path></svg>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return{
      currentChat: null
    }
  },
  methods: {
    SetGlobal() {

    },
    SetFriends() {

    },
    SetClub() {

    }
  }
}

</script>