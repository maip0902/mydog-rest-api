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

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.top_header {
  width: 100%;
  height: 50px;
  background-color:darksalmon
}

.header_app_name {
  float: left;
  margin-top: 10px;
}

.header_app_name a {
  font-weight: bold;
  text-decoration: none;  
  font-size: 24px;
  color: darkslateblue
}

.header_button {
  float: right;
  margin-top: 13px;
}


</style>
