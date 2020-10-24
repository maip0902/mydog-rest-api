<template>
  <div>
    <div>login page</div>
    <form method="post">
      email<input type="email" v-model="email">
      password<input type="text" v-model="password">
      <button @click="login">Login</button>
    </form>
    <router-link to="/signin">Sign In</router-link>
  </div>
</template>

<script>
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
    }
  }
}
</script>

<style scoped>

</style>
