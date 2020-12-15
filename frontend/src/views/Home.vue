<template>
  <div class="main">
    <div class="main-box">
      <img alt="main-img" src="../assets/haru.jpg" class="main-img">
      <div class="main-card">
        <p>
          ステータスコードを入力すると、ステータスに合わせた画像とそのステータスが示す意味が返ってきます！<br>
          ステータスコードを入力して試してみましょう！
        </p>
        <button class="main-btn test-btn">
          <router-link to="/test">MyDogAPIを試す</router-link>
        </button>
      </div>
    </div>
    <StatusImageList/>
    <router-link to="/test">
      <button>
        <strong>Go Test!</strong>
      </button>
    </router-link>
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
