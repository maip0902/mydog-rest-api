<template>
  <div id="app">
    <div class="top-header">
      <router-link to="/"><h1 class="header-logo">MyDogAPI</h1></router-link>
      <div class="header-button">
        <router-link to="/login">
          <button class="main-btn header-btn">
            Admin Login
          </button>
        </router-link>
        <ul class="header-nav">
          <li class="header-nav-itm"><router-link to="/">TOP</router-link></li>
          <li class="header-nav-itm">TEST</li>
          <li class="header-nav-itm">CONTACT</li>
        </ul>
      </div>
    </div>
    <router-view/>
  </div>
</template>

<style lang="scss">
@import "../sass/top.scss";
@import "../sass/login.scss";
</style>
