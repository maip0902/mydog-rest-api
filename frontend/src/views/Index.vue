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
    let params = JSON.stringify({
      token: this.$session.get('jwt')
    })
    axios.post("/api/authUser", params, {headers: {'Content-Type': 'application/json','Accept': 'application/json'}})
      .then((res) => {
        this.$session.set('user', res.data)
        this.isAuthenticated = true
      })
      .catch((err) => {
        this.$session.destroy()
        this.$router.push('/login')
      })
  }
}
</script>

<style scoped>

</style>
