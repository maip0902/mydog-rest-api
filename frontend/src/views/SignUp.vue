<template>
  <div>
    {{ error }}
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
      password: "",
      error: ""
    }
  },
  methods: {
    signup() {
      console.log(this.email)
      let params = JSON.stringify({
        password: this.password,
        email: this.email
      })
      axios.post("http://localhost:3000/signUp", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
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
