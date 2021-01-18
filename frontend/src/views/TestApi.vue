<template>
  <div class="main-content">
      <input type="text" v-model="statusCode" class="main-input">
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
          axios.get("http://localhost/api/code/" + this.statusCode)
            .then((res) => {
                this.description = res.data.Description
                axios.get('http://localhost/api/codeImage/image/' + this.statusCode)
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
.main-content
{
    width: 100%;
    margin-left: auto;
    margin-right: auto;
    margin-top: 5px;
}

.box-image
{
    margin-top: 10px;
    margin-right: 40px;
}

.list-image 
{
  width: 150px;
  height: 150px;
}

.main-input
{
    width: 150px;
    height: 30px;
    border-radius: 8px;
    border:3px solid #FFCCCC;
}

.main-button 
{
    height: 35px;
    border-radius: 8px;
    border:3px solid #FF9966;
    background-color: #FF9966;
    color: #FFFFFF
}

.main-box 
{
    margin-right: 40px;
}

.main-description
{
    font-weight: bold;
    color: #FF9966;
}
</style>