<template>
  <button class="hover:scale-[1.06] hover:text-white w-8 h-8 flex-grow-0 flex justify-center items-center text-[#b3b3b3]">
    <VolumeHigh/>
  </button>
</template>

<script>
import VolumeHigh from '../icon/VolumeHigh.vue';
export default {
  components: {
    VolumeHigh
  }
};
</script>