<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="container">
    <form @submit.prevent="send">
      <div class="input">
        <el-input v-model="inputValue"></el-input>
        <el-button @click="more" id="moreBtn"><el-icon size="1.5em" color="#fff">
            <Files />
          </el-icon></el-button>
        <el-button @click="send" id="sendBtn"><el-icon size="2em" color="#fff">
            <Promotion />
          </el-icon></el-button>
      </div>
    </form>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { ElMessage } from 'element-plus'
import request from '@/axios'

const inputValue = ref('')

const msg = {
  Type: 'text',
  Context: inputValue.value,
  Time: new Date().toLocaleTimeString(),
  From: 'user'
}

const user = computed(() => {
  const userInfo = JSON.parse(localStorage.getItem('chatRoomUserInfo'))
  return userInfo
})

const more = async () => {
  let fileHandles = await window.showOpenFilePicker({
    multiple: true,
  })
  fileHandles.forEach(async (fileHandle) => {
    const file = await fileHandle.getFile()
    // 检测是否是图片
    if (file.type.indexOf('image') === -1) {
      msg.Type = 'file'
      const reader = new FileReader()
      reader.readAsDataURL(file)
      reader.onload = async () => {
        msg.Context = `{"Title": "${file.name}", "Context": "${reader.result}"}`
        msg.Time = `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
        msg.From = user.value
        request.post('/send', msg)
      }
    } else {
      msg.Type = 'image'
      // 获取文件base64内容
      const reader = new FileReader()
      reader.readAsDataURL(file)
      reader.onload = async () => {
        msg.Context = reader.result
        msg.Time = `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
        msg.From = user.value
        request.post('/send', msg)
      }
    }
  })
}

const send = () => {
  if (inputValue.value === '') {
    ElMessage.info('请输入内容')
    return
  }
  msg.Type = 'text'
  msg.Context = inputValue.value
  msg.Time = `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
  msg.From = user.value
  request.post('/send', msg).then(() => {
    inputValue.value = ''
  })
}
</script>

<style scoped>
.container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: #1f1f1f;
  position: absolute;
  bottom: 0;
  left: 0;
  width: calc(100% - 20px);
  z-index: 9999;
  height: 4em;

  form {
    width: 100%;
  }

  .input {
    width: 100%;
    height: 3.5em;
    display: flex;
    justify-content: space-between;
    align-items: center;

    .el-input {
      /* margin-right: 1em; */
      width: 100%;
      height: 2.8em;
    }

    .el-button {
      margin: 0 0 0 0.5em;
      padding: 0;
      border: 0;
      border-radius: 50%;

      &#sendBtn {
        height: 3.5em;
        width: 3.5em;
        min-width: 3.5em;
        background-color: #74b8f7;
        color: #1f1f1f;
      }

      &#moreBtn {
        height: 3em;
        width: 3em;
        min-width: 3em;
        background-color: #2d2d2d;
        color: #fff;
      }
    }
  }
}
</style>