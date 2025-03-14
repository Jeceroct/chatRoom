<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <routeMask />
  <myheader />
  <myfooter />
  <mybody />
</template>

<script setup>
import myheader from '../components/header.vue'
import myfooter from '../components/footer.vue'
import mybody from '../components/body.vue'
import routeMask from '../components/routeMask.vue'
import request from '../axios'
import { Status } from '@/utils/getStatus'
Status.getInstance().stopGetStatus()
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