<template>
  <myheader />
  <myfooter />
  <mybody />
</template>

<script setup>
import myheader from './components/header.vue'
import myfooter from './components/footer.vue'
import mybody from './components/body.vue'
import request from './axios'
// import { onMounted } from 'vue'
// import FastClick from 'fastclick'
// onMounted(async () => {
//   FastClick.attach(document.body)
// })
const getUserInfo = () => {
  request.post('/getUserInfo').then(res => {
    console.log(res)
    localStorage.setItem('chatRoomUserInfo', JSON.stringify(res))
  })
}
getUserInfo()
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
