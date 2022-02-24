<template>
  <button @click="Shuffle" class="hover:scale-[1.06] w-8 h-8 flex-grow-0 flex justify-center items-center">
    <Shuffle/>
  </button>
</template>

<script>
import Shuffle from '../icon/Shuffle.vue';
export default {
  components: {
    Shuffle
  }
}
</script>