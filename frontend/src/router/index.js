import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import TestApi from '../views/TestApi.vue'
import Login from '../views/Login.vue'
import SignUp from "@/views/SignUp";
import Top from "@/views/Top";
import Index from "@/views/Index";
import Edit from "@/views/Edit";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/test',
    name: 'TestApi',
    component: TestApi
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/signup',
    name: 'SignUp',
    component: SignUp
  },
  {
    path: '/top',
    name: 'Top',
    component: Top
  },
  {
    path: '/codeImages',
    name: 'Index',
    component: Index
  },
  {
    path: '/codeImages/:id/edit',
    name: 'Edit',
    component: Edit
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
