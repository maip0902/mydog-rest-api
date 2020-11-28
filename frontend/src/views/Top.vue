<template>
  <div>
    Top
    <router-link to="/codeImages">Index</router-link>
    <button @click="logout">
      logout
    </button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
name: "Top",
  beforeCreate() {
    if (!this.$session.exists()) {
      this.$router.push('/login')
    }
    console.log(this.$session.get('jwt'))
    let params = JSON.stringify({
      token: this.$session.get('jwt')
    })
    axios.post("http://localhost/api/authUser", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
      .then((res) => {
        console.log(res)
        this.$session.set('user', res.data)
        console.log(this.$session.get('user'))
      })
      .catch((err) => {
        console.log(err.response)
    })
  },
  methods: {
    logout() {
      this.$session.destroy()
      this.$router.push('/login')
    }
  }
}
</script>

<style scoped>

</style>
