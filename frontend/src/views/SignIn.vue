<template>
  <div>
    <div>Sign In</div>
    <form method="post">
      email<input type="email" v-model="email">
      password<input type="password" v-model="password">
    </form>
  </div>
</template>

<script>
export default {
  name: "SignIn",
  data() {
    return {
      email: "",
      password: ""
    }
  }
}
</script>

<style scoped>

</style>
