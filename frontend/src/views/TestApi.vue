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
                this.description = res.data.Description
                this.image = require("../assets/" + String(this.statusCode) + ".jpg")
            })
            .catch((err) => {
                console.log(err.response)
                this.image = require("../assets/logo.png")
                this.description = "Sorry, Not Found"
            })
      }
  }
}
</script>

<style scoped>

</style>