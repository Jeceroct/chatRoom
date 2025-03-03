<template>
  <router-view />
</template>

<script setup>
// import request from './axios'
// const getUserInfo = () => {
//   request.post('/getUserInfo').then(res => {
//     console.log(res)
//     localStorage.setItem('chatRoomUserInfo', JSON.stringify(res))
//   })
// }
// getUserInfo()

import { startGetStatus } from '@/utils/getStatus'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
const router = useRouter()

const sleep = (delay) => new Promise((resolve) => setTimeout(resolve, delay))

onMounted(async () => {
  for (; ;) {
    await sleep(1000)
    startGetStatus(router)
  }
})

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
