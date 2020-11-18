<template>
  <div class="main">
    <img alt="Vue logo" src="../assets/haru.jpg" class="top_img">
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

<style scoped>

</style>
