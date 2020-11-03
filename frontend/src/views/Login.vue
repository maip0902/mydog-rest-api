<template>
  <div>
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
      password: ""
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
          console.log(res.data)
          if (res.status == 200) {
            this.$router.push("Top")
          }
        })
    }
  }
}
</script>

<style scoped>

</style>
