<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div class="emojiBoxMask"></div>
  <!-- <textarea class="inputArea" v-model="inputValue" /> -->
  <div id="dropUpload">
    <h2>上传文件</h2>
  </div>
  <EmojiPicker :native="true" :theme="'auto'" :display-recent="true" :disable-skin-tones="true" :hide-search="true"
    :static-texts="{ placeholder: '搜索表情' }" :group-names="emojiGroup" @select="insertEmoji" class="emojiBox" />
  <div class="container">
    <form :submit="send" id="inputForm">
      <div class="input">
        <MyInput v-model="store.inputValue.value" :ref="store.inputRef">
          <template #pre>
            <font-awesome-icon :icon="['fas', 'face-smile']" size="lg" @click="openEmoji" style="cursor: pointer;" />
          </template>
        </MyInput>
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
import { computed, onMounted } from 'vue'
// import { ElMessage } from 'element-plus'
import request from '@/axios'
import RequestType from '@/class/RequestType'
import myMessage from '@/utils/myMessage'
import { onBeforeRouteLeave } from 'vue-router'
import MyInput from './utils/myInput.vue'
import store from '@/store'
import { handlePaste, uploadFile } from '@/utils/myRightMenuFuncs'
import msg from '@/types/msg'
import myDialog from '@/utils/myDialog'

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
  // const inputEle = document.querySelector('.el-input__inner')
  const inputEle = document.querySelector('.myInput textarea')
  inputEle.focus()
  const start = inputEle.selectionStart
  const end = inputEle.selectionEnd
  if (start === undefined || end === undefined) return
  store.inputValue.value = inputEle.value.substring(0, start) + emoji.i + inputEle.value.substring(end)
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
          console.log(e.dataTransfer.files[i])
          if (e.dataTransfer.files[i].type.includes('image')) {
            const reader = new FileReader()
            const file = e.dataTransfer.files[i]
            reader.readAsDataURL(file)
            reader.onload = function (er) {
              // console.log(er)
              myDialog('', 'image', er.target.result, '发送图片', uploadFile, file)
            }
          } else{
            myDialog('文件', 'file', e.dataTransfer.files[i].name, '发送文件', uploadFile, e.dataTransfer.files[i])
          }
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

const send = (e) => {
  if (e) {
    e.preventDefault()
  }
  if (store.inputValue.value === '') {
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
    store.inputValue.value,
    `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
    localStorage.getItem('chatRoomUserInfo') ? JSON.parse(localStorage.getItem('chatRoomUserInfo')) : user.value,
    store.quoteValue
  )
  request.post('/send', msg.getResult()).then((res) => {
    if (res.code != 200) {
      myMessage(res.msg, 'error')
      return
    }
    store.inputValue.value = ''
    store.quoteValue = null
    const emojiBox = document.querySelector('.emojiBox')
    const emojiBoxMask = document.querySelector('.emojiBoxMask')
    emojiBox.classList.remove('show')
    emojiBoxMask.classList.remove('show')
  })
}

onMounted(() => {
  initDropUpload()
  document.addEventListener('paste', handlePaste)
})

onBeforeRouteLeave((to, from, next) => {
  document.removeEventListener('paste', handlePaste)
  next()
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
  height: 3.2em;

  form {
    width: 100%;
  }

  .input {
    width: 100%;
    height: 3.2em;
    display: flex;
    justify-content: space-between;
    align-items: end;

    .myInput {
      /* margin-right: 1em; */
      width: 100%;
      /* height: auto; */
      height: 2.6em;
      border-radius: 1.3em;
      margin-bottom: 0.3em;
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
        margin-bottom: 0.1em;
        background-color: var(--color-theme);
        color: #fff;
      }

      &#moreBtn {
        height: 3em;
        width: 3em;
        min-width: 3em;
        margin-bottom: 0.4em;
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