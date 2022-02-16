<template>
  <div>
    <div class="header">header</div>
    <div class="row">
      <div class="column left">
        <active-users></active-users>
      </div>
      <div class="column middle">
        <router-view/>
        <playback></playback>
      </div>
      <div class="column right"></div>
    </div>
  </div>
</template>

<script>
import ActiveUsers from "./components/social/ActiveUsers.vue"
import Playback from "./components/playback.vue"

export default {
  components: {
    ActiveUsers,
    Playback
  }
}

</script>

<style>
html, body, #app>div {
  height: 100%
}
* {
  box-sizing: border-box;
}

.header {
  background-color: bisque;
  height: 9%;
}

.row {
  height: 91%;
}


.column {
  height: 100%;
}


.column.left {
  width: 17%;
  background-color: aquamarine;
}

.column.right {
  width: 23%;
  background-color: aquamarine;
}
.column.middle {
  width: 60%;
  background-color: darkgray;
}

#app{
  height: 100%;
}

/* width */
::-webkit-scrollbar {
  width: 8px;
}

/* Track */
::-webkit-scrollbar-track {
  background: #f1f1f1;
}

/* Handle */
::-webkit-scrollbar-thumb {
  background: yellowgreen;
}

/* Handle on hover */
::-webkit-scrollbar-thumb:hover {
  background: darkslateblue;
}
</style>