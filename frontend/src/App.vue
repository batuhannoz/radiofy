<template>
  <div class="h-full w-full flex flex-col">
    <Header class="h-[5.3rem] flex-none col"></Header>
    <div class="flex flex-nowrap">
      <active-users class="w-[16.5rem] flex-none"></active-users>
      <div class="flex flex-col flex-grow" style="min-width: 50rem">
        <router-view style="height: calc(100vh - 5.3rem - 92px)"></router-view>
        <player style="height: 92px" class="w-full">player</player>
      </div>
      <Chat class="w-80 flex-none">Chat</Chat>
    </div>
  </div>
</template>

<script>
import ActiveUsers from './components/ActiveUsers.vue'
import Player from './components/Player.vue'
import Chat from  './components/Chat.vue'
import Header from './components/Header.vue'
import VueCookies from "vue-cookies";

export default {
  components: {
    ActiveUsers,
    Player,
    Chat,
    Header
  },
  created() {
    window.onSpotifyWebPlaybackSDKReady = () => {};

    async function waitForSpotifyWebPlaybackSDKToLoad() {
      return new Promise((resolve) => {
        if (window.Spotify) {
          resolve(window.Spotify);
        } else {
          window.onSpotifyWebPlaybackSDKReady = () => {
            resolve(window.Spotify);
          };
        }
      })
    }

    (async () => {
      const { Player } = await waitForSpotifyWebPlaybackSDKToLoad();
      const token = VueCookies.get("access_token");

      // eslint-disable-next-line
      const player = new Player({
        name: "Radiofy",
        getOAuthToken: (cb) => {
          cb(token);
        },
        volume: 0.5
      });
      // Ready
      player.addListener('ready', ({ device_id }) => {
        console.log('Ready with Device ID', device_id);
      });

      // Not Ready
      player.addListener('not_ready', ({ device_id }) => {
        console.log('Device ID has gone offline', device_id);
      });
      player.addListener('initialization_error', ({ message }) => {
        console.error(message);
      });

      player.addListener('authentication_error', ({ message }) => {
        console.error(message);
      });

      player.addListener('account_error', ({ message }) => {
        console.error(message);
      });
      player.connect();

    })()
  }
}
</script>
