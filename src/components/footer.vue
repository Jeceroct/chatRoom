<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="emojiBoxMask"></div>
  <EmojiPicker :native="true" :theme="'dark'" :display-recent="true"
    :static-texts="{ placeholder: '搜索表情', skinTone: '更换肤色' }" :group-names="emojiGroup"
    @select="insertEmoji" class="emojiBox" />
  <div class="container">
    <form @submit.prevent="send">
      <div class="input">
        <el-input v-model="inputValue">
          <template #prefix>
            <font-awesome-icon :icon="['fas', 'face-smile']" size="lg" @click="openEmoji" style="cursor: pointer;" />
          </template>
        </el-input>
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
import RequestType from '@/class/RequestType'

const inputValue = ref('')

const quoteValue = ref(null)

const emojiGroup = {
  "recently_used": "最近使用",
  "smileys_people": "微笑与人物",
  "animals_nature": "动物与自然",
  "food_drink": "食物与饮料",
  "activities": "活动",
  "travel_places": "旅行与地点",
  "objects": "物体",
  "symbols": "符号",
  "flags": "旗帜"
}

// const msg = {
//   Type: 'text',
//   Context: inputValue.value,
//   Time: new Date().toLocaleTimeString(),
//   From: 'user'
// }

const msg = new RequestType(
  RequestType.Type().text,
  '',
  new Date().toLocaleTimeString(),
  RequestType.User(),
  quoteValue
)

const user = computed(() => {
  const userInfo = JSON.parse(localStorage.getItem('chatRoomUserInfo'))
  return userInfo
})

const openEmoji = () => {
  const emojiBox = document.querySelector('.emojiBox')
  const emojiBoxMask = document.querySelector('.emojiBoxMask')
  emojiBox.classList.add('show')
  emojiBoxMask.classList.add('show')
  emojiBoxMask.addEventListener('click', () => {
    emojiBox.classList.remove('show')
    emojiBoxMask.classList.remove('show')
  })
}

const insertEmoji = (emoji) => {
  const inputEle = document.querySelector('.el-input__inner')
  inputEle.focus()
  const start = inputEle.selectionStart
  const end = inputEle.selectionEnd
  if (start === undefined || end === undefined) return
  inputValue.value = inputEle.value.substring(0, start) + emoji.i + inputEle.value.substring(end)
}

const more = async () => {
  let fileHandles = await window.showOpenFilePicker({
    multiple: true,
  })
  fileHandles.forEach(async (fileHandle) => {
    const file = await fileHandle.getFile()
    // 检测是否是图片
    if (file.type.indexOf('image') === -1) {
      // msg.requestType.type = RequestType.Type().file
      const reader = new FileReader()
      reader.readAsDataURL(file)
      reader.onload = async () => {
        // msg.requestType.context = `{"Title": "${file.name}", "Context": "${reader.result}"}`
        // msg.Time = `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
        // msg.From = user.value
        msg.setRequestType(
          RequestType.Type().file,
          RequestType.FileContext(file.name, reader.result),
          `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
          user.value,
          quoteValue
        )
        request.post('/send', msg.getResult())
      }
    } else {
      // msg.Type = 'image'
      // 获取文件base64内容
      const reader = new FileReader()
      reader.readAsDataURL(file)
      reader.onload = async () => {
        // msg.Context = reader.result
        // msg.Time = `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
        // msg.From = user.value
        msg.setRequestType(
          RequestType.Type().image,
          reader.result,
          `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
          user.value,
          quoteValue
        )
        request.post('/send', msg.getResult())
      }
    }
  })
}

const send = () => {
  if (inputValue.value === '') {
    ElMessage.info('请输入内容')
    return
  }
  // msg.Type = 'text'
  // msg.Context = inputValue.value
  // msg.Time = `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`
  // msg.From = user.value
  msg.setRequestType(
    RequestType.Type().text,
    inputValue.value,
    `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
    user.value,
    quoteValue
  )
  request.post('/send', msg.getResult()).then(() => {
    inputValue.value = ''
    quoteValue.value = null
    const emojiBox = document.querySelector('.emojiBox')
    const emojiBoxMask = document.querySelector('.emojiBoxMask')
    emojiBox.classList.remove('show')
    emojiBoxMask.classList.remove('show')
  })
}
</script>

<style scoped>
.emojiBoxMask {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  display: none;
  z-index: 998;

  &.show {
    display: block;
  }
}

.emojiBox {
  position: absolute;
  bottom: -5em;
  left: 0;
  width: 100%;
  visibility: hidden;
  opacity: 0;
  z-index: 999;

  transition: all 0.2s ease-in-out;

  &.show {
    visibility: visible;
    opacity: 1;
    bottom: 5em;
  }
}

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