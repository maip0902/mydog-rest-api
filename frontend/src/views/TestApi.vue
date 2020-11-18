<template>
  <div>
      <input type="text" v-model="statusCode">
      <button>Click!</button>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  name: "TestApi",
  data() {
      return {
          "statusCode": ""
      }
  }
}
</script>

<style scoped>

</style>