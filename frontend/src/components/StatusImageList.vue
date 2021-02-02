<template>
  <div>
    <h2>ステータスコードレスポンス一覧</h2>
    <ul v-for="codeImage in codeImages" v-bind:key="codeImage.id" class="status-list">
      <li class="status-item">
        <img :src="codeImage.Image" class="status-image">
        <p class="status-code">{{ codeImage.Code }}</p>
        <p class="status-dp">{{ codeImage.Description }}</p>
      </li>
      <router-link :to="`/codeImages/${codeImage.ID}/edit`" tag="button" v-if="isAuthenticated" class="main-btn">編集する</router-link>
    </ul>
  </div>

</template>

<script>
import axios from 'axios'

export default {
  name: 'StatusImageList',
  data () {
    return {
      codeImages: [],
      imageData: ""
    }
  },
  created () {
    this.getCode()
  },
  props: {
    'isAuthenticated': {
      type: Boolean,
      default: false,
      required: false
    }
  },
  methods: {
    getCode () {
      axios.get('/api/code')
        .then(res => {
          this.codeImages = res.data
          this.codeImages.map((value) => {
            axios.get('/api/codeImage/image/' + String(value.Code))
              .then((res) => {
                value.Image = res.data.Image == "" ? require("../assets/no-image.jpg") : 'data:image/png;base64,' + res.data.Image
              })
              .catch((e) => {
                console.log(e.response)
                value.Image = require("../assets/no-image.jpg")
              })
          })
        })
    },
  }
}
</script>

<style scoped>
</style>
