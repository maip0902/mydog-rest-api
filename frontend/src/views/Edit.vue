<template>
<div>edit</div>
</template>

<script>
import axios from 'axios'
export default {
name: "Edit",
  created() {
  console.log(this.$route.params.id)
    axios.get('http://localhost:3000/codeImage/' + this.$route.params.id)
      .then((res) => {
        console.log(res.data)
      })
  }
}
</script>

<style scoped>

</style>
