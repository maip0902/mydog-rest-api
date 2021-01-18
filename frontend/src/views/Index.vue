<template>
<div class="main">
  <StatusImageList :isAuthenticated="isAuthenticated"/>
</div>
</template>

<script>
import axios from "axios";
import StatusImageList from "@/components/StatusImageList";

export default {
  name: "Index",
  components: {StatusImageList},
  data() {
    return {
     isAuthenticated: false
    }
  },
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
        console.log(this.$session.get('user'))
        this.$session.set('user', res.data)
        console.log(this.$session.get('user'))
        this.isAuthenticated = true
      })
      .catch((err) => {
        console.log(err.response)
      })
  }
}
</script>

<style scoped>

</style>
