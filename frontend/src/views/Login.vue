<template>
  <div>
    <div>login page</div>
    <form method="post">
      email<input type="email" v-model="username">
      password<input type="text" v-model="password">
      <button @click="login">Login</button>
    </form>
    <router-link to="/signup">Sign Up</router-link>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "Login",
  data () {
    return {
      username: "",
      password: ""
    }
  },
  methods: {
    login() {
      console.log(this.email)
      let params = JSON.stringify({
        username: this.username,
        password: this.password
      })
      axios.post("http://localhost:3000/login", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
        .then( res => {
          console.log(res.data)
        })
    }
  }
}
</script>

<style scoped>

</style>
