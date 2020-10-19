<template>
<div>Code List</div>
  <ul>Code List</ul>
</template>

<script>
import axios from 'axios'

export default {
  name: "StatusImageList",
  created() {
    this.getCode()
  },
  methods: {
    getCode() {
      axios.get("http://localhost:3000/code")
          .then((res =>
            console.log(res.data)
          ))
    }
  }
}
</script>

<style scoped>

</style>