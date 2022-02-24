<template>
  <div class="text-white">
    {{ getAccessToken }}
    {{ getDeviceID}}
  </div>
</template>

<script>
import {mapGetters} from 'vuex';

export default {
  computed: {
    ...mapGetters("auth", ["getAccessToken"]),
    ...mapGetters("player", ["getDeviceID"])
  }
}
</script>
