<template>
  <div>{{token()}}</div>
</template>

<script>
import {setAccessToken, getAvailableDevices} from "@/api/spotify/player";
import axios from "axios";
import VueCookies from 'vue-cookies';
import {mapActions, mapGetters} from 'vuex';

export default {
  computed: {
    ...mapGetters("auth", ["getRadiofyToken", "getAccessToken"]),
  },
  methods: {
    ...mapActions("user", ["refreshUser"]),
    ...mapActions("auth", ["setAccessToken", "setRadiofyToken"]),
    ...mapActions("player", ["setDeviceID", "refreshPlayer"]),
    CompleteAuth(radiofyToken, accessToken) {
      this.setRadiofyToken(radiofyToken)
      this.setAccessToken(accessToken)
      setAccessToken(accessToken)
      this.refreshPlayer()
      this.refreshUser()
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
    },
    token() {
      if (this.getRadiofyToken && this.getAccessToken) {
        this.CompleteAuth(this.getRadiofyToken, this.getAccessToken)
        this.$router.push('/')
      }else {
        const urlSearchParams = new URLSearchParams(window.location.search);
        const params = Object.fromEntries(urlSearchParams.entries());
        if (!params.code) {
          console.log("Unexpected Callback")
          return;
        }
        axios
            .get(
                "http://192.168.1.127:3000/complete_auth",
                {
                  params: {
                    code: params.code,
                    state: params.state,
                    error: params.error
                  }
                }
            )
            .then((res) => {
              VueCookies.set("radiofy_token", res.data.radiofyToken, "1h")
              VueCookies.set('access_token', res.data.accessToken, "1h")
              VueCookies.set('refresh_token', res.data.refreshToken, "1h")
              this.CompleteAuth(res.data.radiofyToken, res.data.accessToken)
            }).then(() => {this.$router.push('/')})
      }
    },
  }
}
</script>