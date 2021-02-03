<template>
  <div>
    <div>管理者画面TOP</div>
    <router-link to="/codeImages" class="main-btn">ステータスコード一覧</router-link>
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
    axios.post("/api/authUser", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json', 'Authentication': this.$session.get('jwt')}})
      .then((res) => {
        console.log(res)
        this.$session.set('user', res.data)
        console.log(this.$session.get('user'))
      })
      .catch((err) => {
        console.log(err.response)
    })
  },
}
</script>

<style scoped>

</style>
