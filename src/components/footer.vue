<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="emojiBoxMask"></div>
  <div id="dropUpload">
    <h2>上传文件</h2>
  </div>
  <EmojiPicker :native="true" :theme="'dark'" :display-recent="true"
    :static-texts="{ placeholder: '搜索表情', skinTone: '更换肤色' }" :group-names="emojiGroup" @select="insertEmoji"
    class="emojiBox" />
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
import { ref, computed, onMounted } from 'vue'
// import { ElMessage } from 'element-plus'
import request from '@/axios'
import RequestType from '@/class/RequestType'
import myMessage from '@/utils/myMessage'

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

// 防抖函数
let timeoutId;
const debounce = (fn, delay) => {
  return (...args) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => fn.apply(this, args), delay);
  };
};


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
  const emojiBox = document.querySelector('.emojiBox')
  const emojiBoxMask = document.querySelector('.emojiBoxMask')
  emojiBox.classList.remove('show')
  emojiBoxMask.classList.remove('show')
}

const initDropUpload = () => {
  const dropUploadEle = document.querySelector('#dropUpload')
  const bodyContainer = document.querySelector('#bodyContainer')
  bodyContainer.addEventListener('dragenter', handlerEvents)
  bodyContainer.addEventListener('dragleave', handlerEvents)
  bodyContainer.addEventListener('dragover', handlerEvents)
  bodyContainer.addEventListener('drop', handlerEvents)

  function handlerEvents(e) {
    e.stopPropagation()
    e.preventDefault()
    switch (e.type) {
      case 'dragenter':
        dropUploadEle.classList.add('show')
        break

      case 'dragover':
        dropUploadEle.classList.add('show')
        break

      case 'dragleave':
        debounce(() => {
          console.log('用户取消上传')
          dropUploadEle.classList.remove('show')
        }, 250).call()
        break

      // 在拖拽区域内松开鼠标（拖放完成/放入文件）
      case 'drop':
        for (let i = 0; i < e.dataTransfer.files.length; i++) {
          uploadFile(e.dataTransfer.files[i])
        }
        dropUploadEle.classList.remove('show')
        break
      default:
        break
    }
  }
}

const more = async () => {
  let fileHandles = await window.showOpenFilePicker({
    multiple: true,
  })
  fileHandles.forEach(async (fileHandle) => {
    const file = await fileHandle.getFile()
    uploadFile(file)
  })
}

const uploadFile = (file) => {
  // 检测是否是图片
  if (file.type.indexOf('image') === -1) {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = async () => {
      msg.setRequestType(
        RequestType.Type().file,
        RequestType.FileContext(file.name, reader.result),
        `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
        user.value,
        quoteValue
      )
      // const elMsg = ElMessage({
      //   message: `正在上传文件: ${file.name}`,
      //   type: 'info',
      //   duration: 0
      // })
      const elMsg = myMessage(`正在上传文件: ${file.name}`, 'info', 0)
      elMsg.load()
      request.post('/send', msg.getResult()).then(() => {
        elMsg.close()
      }).catch(() => {
        // ElMessage.error('文件上传失败')
        myMessage('文件上传失败', 'error')
        elMsg.close()
      })
    }
  } else {
    // 获取文件base64内容
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = async () => {
      msg.setRequestType(
        RequestType.Type().image,
        reader.result,
        `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
        user.value,
        quoteValue
      )
      // const elMsg = ElMessage({
      //   message: `正在上传图片`,
      //   type: 'info',
      //   duration: 0
      // })
      const elMsg = myMessage(`正在上传图片`, 'info', 0)
      elMsg.load()
      request.post('/send', msg.getResult()).then(() => {
        elMsg.close()
      }).catch(() => {
        elMsg.close()
        // ElMessage.error('图片上传失败')
        myMessage('图片上传失败', 'error')
      })
    }
  }
}

const send = () => {
  if (inputValue.value === '') {
    // ElMessage({
    //     message: `请输入消息}`,
    //     type: 'info',
    //     duration: 0
    //   })
    myMessage(`请输入消息`, 'info')
    return
  }
  msg.setRequestType(
    RequestType.Type().text,
    inputValue.value,
    `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
    localStorage.getItem('chatRoomUserInfo') ? JSON.parse(localStorage.getItem('chatRoomUserInfo')) : user.value,
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

onMounted(() => {
  initDropUpload()
})
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

#dropUpload {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  /* height: calc(100% - 10em); */
  background-color: #499dec46;
  display: flex;
  justify-content: center;
  align-items: center;
  color: #fff;
  z-index: 999;

  transition: all 0.2s ease-in-out;
  opacity: 0;
  visibility: hidden;
  pointer-events: none;

  &.show {
    visibility: visible;
    opacity: 1;
  }
}

.container {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  background-color: var(--color-background);
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
        background-color: var(--color-theme);
        color: #fff;
      }

      &#moreBtn {
        height: 3em;
        width: 3em;
        min-width: 3em;
        background-color: var(--color-background-soft);
        color: #fff;
      }
    }
  }
}
</style>