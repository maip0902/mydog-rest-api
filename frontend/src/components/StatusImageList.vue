<template>
  <div>
    <h2>ステータスサンプル</h2>
    <ul v-for="codeImage in codeImages" v-bind:key="codeImage.id" class="status-list">
      <li class="status-item">
        <img :src="imageData" class="status-image">
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
      axios.get('http://localhost/api/code')
        .then(res => {
          this.codeImages = res.data
          this.codeImages.map((value) => {
            this.image(value.Code)
            value.Image = this.image(value.Code)
          })
          console.log(this.codeImages)
        })
    },
    edit () {
      console.log(this.isAuthenticated)
    },
    image(code) {
      axios.get('http://localhost/api/codeImage/image/' + String(code))
        .then((res) => {
          this.imageData = 'data:image/png;base64,' + res.data.Image
          return this.imageData
        })
    }
  }
}
</script>

<style scoped>
</style>
