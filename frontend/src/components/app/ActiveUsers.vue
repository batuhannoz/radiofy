<template>
  <div class="flex flex-col bg-SpotifyBlack text-white overflow-y-auto scroller">
    <div class="bg-SpotifyPlayer h-12 w-auto my-2 mx-[3px] rounded-full justify-center flex flex-row group">
      <img src="https://minimaltoolkit.com/images/randomdata/male/41.jpg" class="rounded-full m-0.5" >
      <div class="ml-2 flex-col flex py-0.5 pl-1 pr-3">
        <div class="font-bold text-[#b3b3b3]" style="max-width: 145px">Jonathan Weddel</div>
        <div class="line-clamp-1 text-xs text-[#b3b3b3]" style="max-width: 145px">Lana Del Rey - Nationel Anthem</div>
      </div>
        <div class="flex rounded-full group-hover:bg-SpotifyGreen my-2 mr-2 h-8 w-8 items-center justify-center">
        <play-icon class="hidden group-hover:block"/>
      </div>
    </div>
  </div>
</template>

<script>
import PlayIcon from "@/assets/icon/Play";

export default {
  components:{
    PlayIcon
  }
}
</script>