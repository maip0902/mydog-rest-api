<template>
  <div>
      <input type="text" v-model="statusCode">
      <button @click="get">Click!</button>
      <div>
          <img :src=image class="top_img"> 
      </div>
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
          image : ""
      }
  },
  methods: {
      get() {
          axios.get("http://localhost:3000/code/" + this.statusCode)
            .then((res) => {
                console.log(res.data)
                this.description = res.data.Description
                console.log(this.statusCode)
                this.image = require("../assets/" + String(this.statusCode) + ".jpg")
                console.log(this.image)
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