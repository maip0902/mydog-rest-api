<template>
  <div>
    {{ error }}
    <div>login page</div>
    email<input type="email" v-model="email">
    password<input type="text" v-model="password">
    <button @click="login">Login</button>
    <router-link to="/signup">Sign Up</router-link>
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
      console.log(this.email)
      let params = JSON.stringify({
        email: this.email,
        password: this.password
      })
      axios.post("http://localhost:3000/signIn", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
        .then( res => {
          this.$router.push("Top")
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
