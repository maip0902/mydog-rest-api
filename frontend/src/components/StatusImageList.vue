<template>
  <div>
    <h2>ステータスサンプル</h2>
    <ul v-for="codeImage in codeImages" v-bind:key="codeImage.id" class="status-list">
      <li class="status-item">
        <img :src=image(codeImage.Code) class="status-image">
        <p class="status-code">{{ codeImage.Code }}</p>
        <p class="status-dp">{{ codeImage.Description }}</p>
        <router-link :to="`/codeImages/${codeImage.ID}/edit`" tag="button" v-if="isAuthenticated">編集する</router-link>
      </li>
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
    }
  },
  created () {
    this.getCode()
    console.log(this.isAuthenticated)
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
          console.log(res.data)
          this.codeImages = res.data
        })
    },
    edit () {
      console.log(this.isAuthenticated)
    },
    image(code) {
      try {
         return require("../assets/" + String(code) + ".jpg") 
      } catch(error) {
          return require("../assets/no-image.jpg")
      }
    }
  }
}
</script>

<style scoped>
</style>
