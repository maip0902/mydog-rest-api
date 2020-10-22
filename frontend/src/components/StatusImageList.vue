<template>
  <div>
    Code List
    <div v-for="codeImage in codeImages" v-bind:key="codeImage.id">
      <dt>{{ codeImage.Code }}</dt>
      <dd>{{ codeImage.Description }}</dd>
    </div>
  </div>

</template>

<script>
import axios from 'axios'

export default {
  name: 'StatusImageList',
  data () {
    return {
      codeImages: []
    }
  },
  created () {
    this.getCode()
  },
  methods: {
    getCode () {
      axios.get('http://localhost:3000/code')
        .then(res => {
          console.log(res.data)
          this.codeImages = res.data
        })
    }
  }
}
</script>

<style scoped>

</style>
