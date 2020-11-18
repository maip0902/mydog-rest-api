<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/haru.jpg">
<!--    <HelloWorld msg="Welcome to Your Vue.js App"/>-->
    <StatusImageList/>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from '@/components/HelloWorld.vue'
import StatusImageList from '@/components/StatusImageList.vue'

export default {
  name: 'Home',
  components: {
    // HelloWorld,
    StatusImageList
  }
}
</script>
