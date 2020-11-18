<template>
  <div>
    Code List
    <div v-for="codeImage in codeImages" v-bind:key="codeImage.id">
      <dt>{{ codeImage.Code }}</dt>
      <dd>{{ codeImage.Description }}</dd>
<!--      <button v-if="isAuthenticated" @click="edit">-->
        <router-link :to="`/codeImages/${codeImage.ID}/edit`" tag="button" v-if="isAuthenticated">編集する</router-link>
<!--      </button>-->
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

    }
  }
}
</script>

<style scoped>

</style>
