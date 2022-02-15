<template>
  <div class="LoginPage">
    <login-button></login-button>
  </div>
</template>

<script>
import LoginButton from "@/components/login/LoginButton";

export default {
  components: {
    LoginButton
  },
  data() {
    return {
      page: LoginButton
    }
  }
}
</script>

<style>
.LoginPage{
  position: relative;
  height: 89%;

}

</style>