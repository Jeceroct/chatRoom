<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <routeMask />
  <div class="cr">
    <h1>无法连接到聊天室</h1>
    <p class="tips">这可能是你的网络出现了问题，也有可能是此聊天室已关闭</p>
    <button class="quitBtn" @click="quit">换一个聊天室</button>
    <button class="reConnBtn" @click="reConn">重新连接</button>
  </div>
</template>

<script setup>
import routeMask from '../components/routeMask.vue'
import request from '../axios'
import { useRouter } from 'vue-router'
const router = useRouter()

const quit = () => {
  const routeMaskEle = document.querySelector('#routeMask')
  routeMaskEle.classList.add('leave')
  routeMaskEle.classList.remove('waiting')
  setTimeout(() => {
    router.push('/address')
  }, 200)
}

const reConn = () => {
  request.post('/reConn', {}, {}).then(res => {
    if (res.code == '200') {
      const routeMaskEle = document.querySelector('#routeMask')
      routeMaskEle.classList.add('leave')
      routeMaskEle.classList.remove('waiting')
      setTimeout(() => {
        router.push('/')
      }, 200)
    }
  })
}
</script>

<style scoped>
.cr {
  height: 100vh;
  width: 100vw;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  &.tips {
    margin-top: 1em;
    color: rgb(175, 175, 175);
    font-size: 0.8em;
  }
}

button {
  width: 80%;
  height: 3em;
  border-radius: 2em;
  margin: 0.8em 0;
  background-color: rgb(98, 149, 110);
  border: 2px solid rgb(98, 149, 110);
  color: white;
  cursor: pointer;

  &.quitBtn {
    margin-top: 3em;
    background-color: rgba(214, 255, 224, 0.496);
  }
}
</style>