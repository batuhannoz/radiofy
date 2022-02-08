<template>
  <div>
    <va-slider v-model="changeVolume" :step="10"/>
    <div>{{value}}</div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data () {
    return {
      changeVolume: 45,
    }
  },
  watch: {
    changeVolume(value) {
      axios.get("http://localhost:8080/volume", {params: {volume: value}} )
    }
  }
}
</script>