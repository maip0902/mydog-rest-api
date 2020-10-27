<template>
  <div>
    <div>Sign up</div>
      email<input type="email" v-model="email">
      password<input type="password" v-model="password">
      <button @click="signup">register</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "SignUp",
  data() {
    return {
      email: "",
      password: ""
    }
  },
  methods: {
    signup() {
      axios.post("http://localhost:3000/signUp", {"email": this.email, "password": this.password}, {headers: {'Content-Type': 'application/x-www-form-urlencoded'}})
        .then( res => {
          console.log(res.data)

        })
    }
  }
}
</script>

<style scoped>

</style>
