<template>
<div>hello</div>
  <input type="text" name="code" v-model="code">
  <button @click="getImage">Get Image!</button>
  <div>
    <div>{{ description }}</div>
    <img :src="imageUrl">
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "StatusImage",
  data () {
    return {
      code: "",
      imageUrl: "",
      description: ""
    }
  },
  methods: {
    getImage() {
      axios.get("http://localhost:3000/code/" + this.code)
          .then(res => (
              console.log(res.data.Image),
              this.imageUrl = "http://localhost:9000/minio/" + res.data.Image,
              this.description = res.data.Description

          ))
          .catch(err => console.log(err))
    }
  }
}
</script>

<style scoped>

</style>