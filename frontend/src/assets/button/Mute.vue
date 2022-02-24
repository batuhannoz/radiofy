<template>
  <button class="hover:scale-[1.06] hover:text-white w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
    <Mute/>
  </button>
</template>

<script>
import Mute from '../icon/Mute.vue';
export default {
  components: {
    Mute
  }
}
</script>