<template>
  <div>{{token()}}</div>
</template>

<script>
import {setAccessToken, getAvailableDevices} from "@/api/spotify/player";
import axios from "axios";
import VueCookies from 'vue-cookies';
import {mapActions} from 'vuex';

export default {
  methods: {
    ...mapActions("auth", ["setAccessToken"]),
    ...mapActions("player", ["setDeviceID", "refreshPlayer"]),
    token() {
      console.log("hello")
      const urlSearchParams = new URLSearchParams(window.location.search);
      const params = Object.fromEntries(urlSearchParams.entries());
      if (!params.code) {
        console.log("Unexpected Callback")
        return;
      }
      console.log(params)
      axios
          .get(
              "http://localhost:3000/complete_auth",
              {
                params: {
                  code: params.code,
                  state: params.state,
                  error: params.error
                }
              }
          )
          .then((res) => {
            VueCookies.set('access_token', res.data.accessToken, "1h")
            VueCookies.set('refresh_token', res.data.refreshToken, "1h")
            this.setAccessToken(res.data.accessToken)
            setAccessToken(res.data.accessToken)
            this.refreshPlayer()
            getAvailableDevices().then((res) => {
              this.setDeviceID(res.devices[0].id)
            })
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
                volume: 0.3
              });
              // Ready
              player.addListener('ready', ({ device_id }) => {
                console.log('Ready with Device ID', device_id);
                this.setDeviceID(device_id)
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

          }).then(() => {this.$router.push('/')})
    },
  }
}
</script>