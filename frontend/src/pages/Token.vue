<script>
  import VueCookies from 'vue-cookies';
  import axios from "axios";
  import queryString from "query-string";

  export default {
    beforeCreate() {
      const urlSearchParams = new URLSearchParams(window.location.search);
      const params = Object.fromEntries(urlSearchParams.entries());
      if (!params.code) {
        console.log("token alınamadı")
        return;
      }
      axios
          .post(
              "https://accounts.spotify.com/api/token",
              queryString.stringify({
                code: params.code,
                redirect_uri: "http://localhost:8080/token",
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
            VueCookies.set('access_token' , res.data.access_token, "30min")
          }).then(setTimeout(() => {this.$router.push('/')},1000))
    }
  }
</script>