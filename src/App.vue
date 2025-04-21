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

import { Status } from '@/utils/getStatus'
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import myRightClickMenu from './utils/myRightClickMenu'
import myMenuItem from './utils/myMenuItem'
import { handlePaste, refreshPage } from './utils/myRightMenuFuncs'
import store from './store'

var clipboardItem

// 监听右键点击
const listenRightClick = async (e) => {
  e.preventDefault()

  // 监听剪切板内容
  if (store.currentPage.value == 'start') {
    clipboardItem = await navigator.clipboard.read().then((clipboardItems) => { return clipboardItems[0] })
    let item = void 0
    // console.log(clipboardItem)
    if (clipboardItem.types.length > 0 && clipboardItem.types[0].includes('text')) {
      item = new myMenuItem('粘贴', handlePaste, clipboardItem)
    } else if (clipboardItem.types.length > 0 && clipboardItem.types[0].includes('image')) {
      item = new myMenuItem('发送图片', handlePaste, clipboardItem)
    } else if (clipboardItem.types.length > 0) {
      item = new myMenuItem('发送文件', handlePaste, clipboardItem)
    }
    if (item != void 0) {
      store.rightClickMenu.addItem(item)
    }
  }

  // 加入刷新页面选项
  store.rightClickMenu.addItem(new myMenuItem('刷新页面(慎用)', refreshPage))

  store.rightClickMenu.show(e.clientX, e.clientY)
  store.rightClickMenu.removeAllItems()

}

onMounted(async () => {
  // 监听右键点击
  document.addEventListener('contextmenu', listenRightClick)
  // 监听点击事件关闭右键菜单
  document.addEventListener('click', myRightClickMenu.close)
  const menu = new myRightClickMenu()
  store.rightClickMenu = menu
  Status.getInstance(useRouter())
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
