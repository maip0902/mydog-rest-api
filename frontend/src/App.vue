<template>
  <div id="app">
    <div class="top-header">
      <router-link to="/"><h1 class="header-logo">MyDogAPI</h1></router-link>
      <div class="header-button">
        <router-link v-if="!isAuthenticated" to="/login">
          <button class="main-btn header-btn">
            Admin Login
          </button>
        </router-link>
        <button v-if="isAuthenticated" class="main-btn header-btn" @click="logout">
            Logout
        </button>
        <ul class="header-nav" v-if="!isAuthenticated">
          <li class="header-nav-itm"><router-link to="/">TOP</router-link></li>
          <li class="header-nav-itm"><router-link to="/test">TEST</router-link></li>
          <li class="header-nav-itm">CONTACT</li>
        </ul>
        <ul class="header-nav" v-if="isAuthenticated">
          <li class="header-nav-itm"><router-link to="/codeImages">TOP</router-link></li>
          <li class="header-nav-itm">CONTACT</li>
        </ul>
      </div>
    </div>
    <router-view/>
  </div>
</template>

<script>
import axios from 'axios'

export default {
  data () {
    return {
      isAuthenticated: this.$session.has('jwt')
    }
  },
  methods: {
    logout() {
        this.$session.destroy();
        location.reload()
        this.$router.push('Login')
    }
  }
}
</script>

<style lang="scss">
@import "../sass/common.scss";
@import "../sass/top.scss";
@import "../sass/login.scss";
@import "../sass/edit.scss";
@import "../sass/test.scss";
</style>
