<template>
  <div>
    Code List
    <div v-for="codeImage in codeImages" v-bind:key="codeImage.id">
      <div>{{ codeImage.Code }}</div>
      <img :src=image(codeImage.Code) class="list-image">
      <div>{{ codeImage.Description }}</div>
        <router-link :to="`/codeImages/${codeImage.ID}/edit`" tag="button" v-if="isAuthenticated">編集する</router-link>
    </div>
    <router-link to="/test">
      <button>
        <strong>Go Test!</strong>
      </button>
    </router-link>
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
      axios.get('http://localhost:3000/code')
        .then(res => {
          console.log(res.data)
          this.codeImages = res.data
        })
    },
    edit () {
      console.log(this.isAuthenticated)
    },
    image(code) {
        return require("../assets/" + String(code) + ".jpg") 
    }
  }
}
</script>

<style scoped>
.list-image 
{
  width: 150px;
  height: 150px;
}
</style>
