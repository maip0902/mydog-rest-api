<template>
<div>hello</div>
  <input type="text" name="code" v-model="code">
  <button @click="getImage">Get Image!</button>
  
  <img :src="imageUrl">
</template>

<script>
import axios from 'axios'

export default {
  name: "StatusImage",
  data () {
    return {
      code: "",
      imageUrl: ""
    }
  },
  methods: {
    getImage() {
      axios.get("http://localhost:3000/code/" + this.code)
          .then(res => (
              console.log(res.data.Image),
              this.imageUrl = require('@/assets/' + res.data.Image + '.jpeg')
          ))
          .catch(err => console.log(err))
    }
  }
}
</script>

<style scoped>

</style>