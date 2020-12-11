<template>
  <div id="app">
    <div class="top_header">
      <div class="header_app_name">
        <router-link to="/"><strong>MyDogAPI</strong></router-link>
      </div>
      <div class="header_button">
        <router-link to="/login">Admin Login</router-link>
      </div>
    </div>
    <router-view/>
  </div>
</template>

<style lang="scss">
@import "../sass/top.scss";
</style>
