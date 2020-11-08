<template>
  <div>
    <div>code : {{ code }}</div>
    <div>
      Description <input type="textarea" v-model="description">
    </div>
    <button>編集する</button>
  </div>
</template>

<script>
import axios from 'axios'
export default {
name: "Edit",
  data () {
    return {
     code: "",
     description: "",
     image: ""
    }
  },
  created() {
  console.log(this.$route.params.id)
    axios.get('http://localhost:3000/codeImage/' + this.$route.params.id)
      .then((res) => {
        console.log(res.data)
        this.code = res.data.Code
        this.description = res.data.Description
        this.image = res.data.Image
      })
  }
}
</script>

<style scoped>

</style>
