<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="emojiBoxMask"></div>
  <textarea class="inputArea" v-model="inputValue" />
  <div id="dropUpload">
    <h2>上传文件</h2>
  </div>
  <EmojiPicker :native="true" :theme="'auto'" :display-recent="true" :disable-skin-tones="true" :hide-search="true"
    :static-texts="{ placeholder: '搜索表情' }" :group-names="emojiGroup" @select="insertEmoji" class="emojiBox" />
  <div class="container">
    <form @submit.prevent="send">
      <div class="input">
        <el-input v-model="inputValue" ref="inputRef">
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
import { ref, computed, onMounted, watch } from 'vue'
// import { ElMessage } from 'element-plus'
import request from '@/axios'
import RequestType from '@/class/RequestType'
import myMessage from '@/utils/myMessage'
import myDialog from '@/utils/myDialog'

const inputRef = ref(null)

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
  request.post('/send', msg.getResult()).then((res) => {
    if (res.code != 200) {
      myMessage(res.msg, 'error')
      return
    }
    inputValue.value = ''
    quoteValue.value = null
    const emojiBox = document.querySelector('.emojiBox')
    const emojiBoxMask = document.querySelector('.emojiBoxMask')
    emojiBox.classList.remove('show')
    emojiBoxMask.classList.remove('show')
  })
}

// 监听粘贴事件
const handlePaste = (e) => {
  // 检查粘贴的内容是否为图片
  e.preventDefault()
  console.log(e.clipboardData)
  if (e.clipboardData && e.clipboardData.files && e.clipboardData.files.length > 0) {
    const file = e.clipboardData.files[0]
    const reader = new FileReader()
    reader.readAsDataURL(file)
    if (file.type.indexOf('image') === -1) {
      // 粘贴文件
      reader.onload = () => {
        myDialog('文件', 'file', file.name, '发送文件', uploadFile, file)
      }
    } else {
      // 粘贴图片
      reader.onload = () => {
        myDialog('', 'image', reader.result, '发送图片', uploadFile, file)
      }
    }
  } else {
    // 粘贴的内容不是图片，执行默认粘贴操作
    const pastedText = e.clipboardData.getData('text/plain')
    inputValue.value = pastedText
    inputRef.value.focus()
  }
}

var inputArea
// var inputEle
// 监听输入长度
const handleInput = () => {
  if (inputValue.value.length > 20) {
    inputArea.classList.add('show')
  } else if (inputArea.classList.contains('show')) {
    inputArea.classList.remove('show')
    inputRef.value.focus()
  }
}

onMounted(() => {
  initDropUpload()
  inputArea = document.querySelector('.inputArea')
  // inputEle = document.querySelector('.input .el-input')
  watch(inputValue, handleInput)
  document.addEventListener('paste', handlePaste)
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

.inputArea {
  background-color: #fff;
  position: absolute;
  bottom: 0;
  left: 50%;
  transform: translateX(-50%);
  border-radius: 1.4em;
  z-index: 900;
  opacity: 0;
  transition: all 0.2s ease;
  pointer-events: none;
  font-size: 1.2em;
  border: 0;
  width: 80vw;
  height: 150px;
  padding: 10px 25px;

  font-weight: 400;
  letter-spacing: 0.5px;

  white-space: pre-wrap;
  word-wrap: break-word;
  word-break: break-all;
  overflow-y: scroll;

  &.show {
    opacity: 1;
    pointer-events: all;
    overflow-y: scroll;
    bottom: 6em;
  }

  &::-webkit-scrollbar {
    display: none;
  }
}
</style>