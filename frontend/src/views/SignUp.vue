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
      // let params = new URLSearchParams();
      // params.append("email", this.email);
      // params.append( "password", this.password)
      let data = JSON.stringify({
        password: this.password,
        email: this.email
      })
      axios.post("http://localhost:3000/signUp", data, {headers: {'Content-Type': 'application/json', 'Accept': 'application/json'}})
        .then( res => {
          console.log(res.data)

        })
    }
  }
}
</script>

<style scoped>

</style>
