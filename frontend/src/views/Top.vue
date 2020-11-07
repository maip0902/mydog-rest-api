<template>
  <div>
    Top
    <button @click="logout">
      logout
    </button>
  </div>
</template>

<script>
export default {
name: "Top",
  beforeCreate() {
  console.log(this.$session.exists())
    if (!this.$session.exists()) {
      this.$router.push('/signUp')
    }
  },
  methods: {
    logout() {
      this.$session.destroy()
      this.$router.push('/signUp')
    }
  }
}
</script>

<style scoped>

</style>
