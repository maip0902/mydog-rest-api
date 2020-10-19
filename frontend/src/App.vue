<template>
<!--  <img alt="Vue logo" src="./assets/logo.png">-->
<!--  <HelloWorld msg="Welcome to Your Vue.js App"/>-->
  <StatusImage/>
  <StatusImageList/>
</template>

<script>
import StatusImage from './components/StatusImage.vue'
import StatusImageList from "@/components/StatusImageList";

export default {
  name: 'App',
  components: {
    StatusImageList,
    StatusImage
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
