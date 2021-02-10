<template>
<div>
  <div v-if="isSuccess">
    本登録が完了しました
  </div>
  <div v-if="!isSuccess">
    本登録が完了できませんでした。再度登録をお願いします。
  </div>
</div>  
</template>

<script>
import axios from 'axios';

export default {
  name: "Registered",
  data() {
    return {
      "isSuccess": true
    }
  },
  created() {
      console.log(this.$route.query)
    axios.get("/api/email?verify_token=" + this.$route.query.verify_token)
      .then((res) => {
        console.log(res)
        this.isSuccess = true
      })
      .catch((e) => {
        console.log(e)
        this.isSuccess = false
      })
  }
}
</script>

<style scoped>

</style>