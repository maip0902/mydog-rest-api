<template>
  <div>
    {{ error }}
    <div class="login-input">
      <label for="email" class="login-label">email</label>
      <input id="email" type="email" class="login-field" v-model="email">
    </div>
    
    <div class="login-input">
      <label for="password" class="login-label">password</label>
      <input id="password" type="password" class="login-field" v-model="password">
    </div>
    
    <div class="btn-field">
      <button id="loginBtn" class="main-btn login-btn" @click="login">ログイン</button>
      <button class="main-btn register-btn" onclick="location.href='/signup'">新規会員登録</button>
    </div>  
    
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "Login",
  data () {
    return {
      email: "",
      password: "",
      error: ""
    }
  },
  methods: {
    login() {
      let params = JSON.stringify({
        email: this.email,
        password: this.password
      })
      axios.post("/api/signIn", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
        .then( res => {
          console.log(res)
          this.$session.set('jwt', res.data.Token)
          // this.$router.push("Top")
          // location.reload()
        })
        .catch((error) => {
          this.error = error.response.data.Error
        })
    }
  }
}
</script>

<style scoped>

</style>
