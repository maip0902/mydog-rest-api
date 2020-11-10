<template>
  <div>
    <div>code : {{ code }}</div>
    <div>
      Description <input type="textarea" v-model="description">
      画像 <input type="file" @change="upload">
    </div>
    <button @click="update">編集する</button>
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
  },
  methods: {
    upload(event) {
      console.log(event)
    },
    update() {
      let params = JSON.stringify({
        id: this.$route.params.id,
        description: this.description
      })
      axios.post('http://localhost:3000/codeImage/' + this.$route.params.id, params,{"headers": {"Content-Type": "application/json", "Accept": "application/json"}})
        .then((res) => {
          console.log(res)
        })
        .catch((err) => {
          console.log(err.response)
        })
    }
  }
}
</script>

<style scoped>

</style>
