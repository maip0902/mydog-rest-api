<template>
<div>hello</div>
  <input type="text" name="code" v-model="code">
  <button @click="getImage">Get Image!</button>
</template>

<script>
import axios from 'axios'

export default {
  name: "StatusImage",
  data () {
    return {
      code: ""
    }
  },
  methods: {
    getImage() {
      console.log(this.code)
      axios.get("http://localhost:9000/code/" + this.code)
          .then(res => (console.log(res.data)))
          .catch(err => console.log(err))
    }
  }
}
</script>

<style scoped>

</style>