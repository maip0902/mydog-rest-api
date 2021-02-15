<template>
  <div class="main-content">
      <input type="text" v-model="statusCode" placeholder="200" class="main-input">
      <button @click="get" class="main-button">Click!</button>
      <div class="box-image">
          <img :src="image" class="list-image"> 
      </div>
    <div class="main-box">
      <div class="main-description">{{description}}</div>
    </div>
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
          "image" : require("../assets/no-image.jpg")
      }
  },
  methods: {
      get() {
          axios.get("/api/code/" + this.statusCode)
            .then((res) => {
                this.description = res.data.Description
                axios.get('/api/codeImage/image/' + this.statusCode)
                    .then((res) => {
                        this.image = res.data.Image == "" ? require("../assets/no-image.jpg") : 'data:image/png;base64,' + res.data.Image
                    })
            })
            .catch((err) => {
                console.log(err.response)
                this.image = require("../assets/no-image.jpg")
                this.description = "Sorry, Not Found"
            })
      }
  }
}
</script>

<style scoped>

</style>