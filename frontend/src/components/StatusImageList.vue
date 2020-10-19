<template>
<div>Code List</div>
  <ul>Code List</ul>
  <li v-for="codeImage in codeImages" v-bind:key="codeImage.id">
    {{ codeImage.Description }}
  </li>
</template>

<script>
import axios from 'axios'

export default {
  name: "StatusImageList",
  data() {
    return {
      codeImages: []
    }
  },
  created() {
    this.getCode()
  },
  methods: {
    getCode() {
      axios.get("http://localhost:3000/code")
          .then(res => (
            this.codeImages = res.data,
                console.log(res.data)
          ))
    }
  }
}
</script>

<style scoped>

</style>