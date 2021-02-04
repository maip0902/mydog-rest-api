<template>
  <div>
    <p class="text-danger">{{ error }}</p>
    <div class="login-input">
      <label for="email" class="login-label">メールアドレス</label>
      <input id="email" type="email" class="login-field" v-model="email">
    </div>
    
    <div class="login-input">
      <label for="password" class="login-label">パスワード</label>
      <input id="password" type="password" class="login-field" v-model="password">
    </div>

    <div class="login-input">
      <label for="password_confirm" class="login-label">パスワード(確認用)</label>
      <input id="password_confirm" type="password" class="login-field" v-model="passwordConfirm">
    </div>
    
    <div class="btn-field">
      <button id="loginBtn" class="main-btn login-btn" @click="signup">新規会員登録</button>
    </div>  
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "SignUp",
  data() {
    return {
      email: "",
      password: "",
      passwordConfirm: "",
      error: ""
    }
  },
  methods: {
    signup() {
      if(this.password !== this.passwordConfirm) {
        this.error = "パスワードは確認用と一致しません"
        return
      }
      let params = JSON.stringify({
        password: this.password,
        email: this.email
      })
      axios.post("/api/signUp", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
        .then( res => {
          try {
            this.$session.start()
            this.$session.set('jwt', res.data.Token)
            this.$router.push("Top")
            location.reload()
          } catch (e) {
            console.log(e)
          }
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
