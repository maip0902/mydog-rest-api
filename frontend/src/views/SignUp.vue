<template>
  <div>
    <div>Sign up</div>
    <form method="post">
      email<input type="email" v-model="email">
      password<input type="password" v-model="password">
      <button @click="signup">register</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "SignIn",
  data() {
    return {
      email: "",
      password: ""
    }
  },
  methods: {
    signup() {
      axios.post("/signup", {"email": this.email, "password": this.password})
        .then( res => {
          console.log(res.data)
        })
    }
  }
}
</script>

<style scoped>

</style>
