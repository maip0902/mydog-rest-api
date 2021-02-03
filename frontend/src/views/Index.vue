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
    axios.post("/api/authUser", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
      .then((res) => {
        console.log(this.$session.get('user'))
        this.$session.set('user', res.data)
        console.log(this.$session.get('user'))
        this.isAuthenticated = true
      })
      .catch((err) => {
        if(err.response.status === 401) {
          this.$session.destroy()
          this.$router.push('/login')
        }
      })
  }
}
</script>

<style scoped>

</style>
