<template>
  <div>{{token()}}</div>
</template>

<script>
import {setAccessToken} from "@/api/spotify/player";
import axios from "axios";
import VueCookies from 'vue-cookies';
import queryString from 'query-string';
import {mapActions} from 'vuex';

export default {
  methods: {
    ...mapActions("auth", ["setAccessToken"]),
    ...mapActions("player", ["setDeviceID", "refreshPlayer"]),
    token() {
      const urlSearchParams = new URLSearchParams(window.location.search);
      const params = Object.fromEntries(urlSearchParams.entries());
      if (!params.code) {
        console.log("Unexpected Callback")
        return;
      }
      axios
          .post(
              "https://accounts.spotify.com/api/token",
              queryString.stringify({
                code: params.code,
                redirect_uri: "http://localhost:8080/callback",
                grant_type: "authorization_code",
              }),
              {
                headers: {
                  "Content-Type": "application/x-www-form-urlencoded",
                  Authorization:
                      "Basic " +
                      new Buffer(
                          "d97ea09987ab46c2bb0b4fa4eaae55e1" +
                          ":" +
                          "cda1b2e945234060a43f0791f4839397"
                      ).toString("base64"),
                },
              }
          )
          .then((res) => {
            VueCookies.set('access_token', res.data.access_token, "1h")
            this.setAccessToken(res.data.access_token)
            setAccessToken(res.data.access_token)
            this.refreshPlayer()

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