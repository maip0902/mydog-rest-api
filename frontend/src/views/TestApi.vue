<template>
  <div>
      <input type="text" v-model="statusCode">
      <button @click="get">Click!</button>
      <dt>{{description}}</dt>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "TestApi",
  data() {
      return {
          "statusCode": "",
          "description": "",
      }
  },
  methods: {
      get() {
          axios.get("http://localhost:3000/code/" + this.statusCode)
            .then((res) => {
                console.log(res.data)
                this.description = res.data.Description
            })
            .catch((err) => {
                console.log(err.response)
            })
      }
  }
}
</script>

<style scoped>

</style>