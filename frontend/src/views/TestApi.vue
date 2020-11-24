<template>
  <div class="main-content">
      <input type="text" v-model="statusCode" class="main-input">
      <button @click="get" class="main-button">Click!</button>
      <div class="box-image">
          <img :src=image class="list-image"> 
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
.main-content
{
    width: 100%;
    margin-left: auto;
    margin-right: auto;
    margin-top: 5px;
}

.box-image
{
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