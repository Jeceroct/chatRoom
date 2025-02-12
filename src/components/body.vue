<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <div id="bodyContainer">
    <div class="msgBox">

    </div>
  </div>
  <div class="newMsg">
    您有新消息
    <span class="num">{{ newMsgNum }}</span>
  </div>
  <div class="scrollToBtm">
    <el-icon size="1.8em">
      <DArrowRight />
    </el-icon>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import request from '@/axios'

// const fs = require('fs')
// const path = require('path')
// const dataPath = path.join(__dirname, 'data.json')
// const config = fs.readFileSync(dataPath, 'utf-8')
// const datas = JSON.parse(config)

const user = JSON.parse(localStorage.getItem('chatRoomUserInfo'))

const newMsgNum = ref(0)

const sleep = (delay) => new Promise((resolve) => setTimeout(resolve, delay))
// const msgsData = ref("")

const addMsgHTML = (msg, msgBoxEle) => {
  var newHtml = ""
  if (msgBoxEle.hasChildNodes() && msgBoxEle.lastChild.getAttribute("userId") == msg.From.Id) {
    if (msg.From.Id == user.Id) {
      newHtml = `
          <div class="msg myself" userId="${msg.From.Id}">
            <div class="textBox">
              <div class="time">${msg.Time}</div>
              <div class="text">`+ msg.Context + `</div>
            </div>
          </div>`
    } else {
      newHtml = `
            <div class="msg" userId="${msg.From.Id}">
              <div class="textBox">
                <div class="text">`+ msg.Context + `</div>
                <div class="time">${msg.Time}</div>
              </div>
            </div>`
    }
  } else {
    if (msg.From.Id == user.Id) {
      newHtml = `
          <div class="msg first myself" userId="${msg.From.Id}">
            <div class="user">
              <div class="level">Lv${msg.From.Level}</div>
              <div class="title" style="background-color:${msg.From.TitleColor}">${msg.From.Title}</div>
              <div class="name">${msg.From.Name}</div>
            </div>
            <div class="textBox">
              <div class="time">${msg.Time}</div>
              <div class="text">`+ msg.Context + `</div>
            </div>
          </div>`
    } else {
      newHtml = `
            <div class="msg first" userId="${msg.From.Id}">
              <div class="user">
                <div class="name">${msg.From.Name}</div>
                <div class="title" style="background-color:${msg.From.TitleColor}">${msg.From.Title}</div>
                <div class="level">Lv${msg.From.Level}</div>
              </div>
              <div class="textBox">
                <div class="text">`+ msg.Context + `</div>
                <div class="time">${msg.Time}</div>
              </div>
            </div>`
    }
  }
  if (msg.Type == 'image') {
    newHtml = newHtml.replace(`<div class="text">` + msg.Context + `</div>`, `<img src="${process.env.VUE_APP_API_ADDR}${msg.Context}">`)
    msgBoxEle.innerHTML += newHtml
  } else
  if (msg.Type == 'file') {
    newHtml = newHtml.replace(`<div class="text">` + msg.Context + `</div>`, `<div class="text file">${msg.Context}</div>`)
    msgBoxEle.innerHTML += newHtml
    msgBoxEle.lastChild.addEventListener('click', () => {
      request.post('/download', msg)
    })
  } else {
    msgBoxEle.innerHTML += newHtml
  }
  // msgsData.value += newHtml
}

onMounted(async () => {
  const container = document.querySelector('#bodyContainer')
  // 加载历史消息(只加载最后100条)
  const msgBoxEle = document.querySelector('.msgBox')
  request.post('/getHistory').then(res => {
    if (res) {
      if (res.length > 100) {
        res = res.slice(res.length - 100, res.length)
      }
      res.forEach(i => {
        addMsgHTML(i, msgBoxEle)
      });
      container.style.scrollBehavior = 'unset'
      container.scrollTop = container.scrollHeight
      container.style.scrollBehavior = 'smooth'
    }
  })

  // 定义滚动到底部按钮的位置
  const scrollToBtmEle = document.querySelector('.scrollToBtm')
  const sendBtnEle = document.querySelector('#sendBtn')
  scrollToBtmEle.style.left = sendBtnEle.offsetLeft + 'px'

  // 定义新消息按钮的位置
  const newMsgElement = document.querySelector('.newMsg')
  newMsgElement.style.left = container.clientWidth / 2 - newMsgElement.clientWidth / 2 + 'px'

  // 监听container滚动事件，滚动到底部时，隐藏newMsg和scrollToBtm
  container.addEventListener('scroll', () => {
    const isScrolledToBottom = container.scrollTop + container.clientHeight >= container.scrollHeight - 100
    if (isScrolledToBottom) {
      newMsgElement.classList.remove('show')
      scrollToBtmEle.classList.remove('show')
      newMsgNum.value = 0
    } else if (!newMsgElement.classList.contains('show')) {
      scrollToBtmEle.classList.add('show')
    }
  })

  // 监听newMsg点击事件，滚动到底部
  newMsgElement.addEventListener('click', () => {
    container.scrollTop = container.scrollHeight
    newMsgElement.classList.remove('show')
    newMsgNum.value = 0
  })

  // 监听scrollToBtm点击事件，滚动到底部
  scrollToBtmEle.addEventListener('click', () => {
    container.scrollTop = container.scrollHeight
    scrollToBtmEle.classList.remove('show')
  })

  // 监听新消息
  for (; ;) {
    const error = ref(null)
    const res = await request.post('/get').catch(err => error.value = err)
    if (error.value) {
      await sleep(5000)
    } else {
      console.log(res.message)
      res.message.forEach(i => {
        addMsgHTML(i, msgBoxEle)
      });

      // 判断container内部是否已经滚动到最底部
      const isScrolledToBottom = container.scrollTop + container.clientHeight >= container.scrollHeight - 100
      if (isScrolledToBottom) {
        container.scrollTop = container.scrollHeight - container.clientHeight
      } else {
        const newMsg = document.querySelector('.newMsg')
        newMsg.classList.add('show')
        scrollToBtmEle.classList.remove('show')
        newMsgNum.value += res.message.length
      }
    }
  }
})
</script>

<style>
#bodyContainer {
  display: flex;
  flex-direction: column;
  justify-content: start;
  align-items: start;
  position: absolute;
  width: 100%;
  top: 5em;

  height: calc(100vh - 10em);

  overflow-y: scroll;
  scroll-behavior: smooth;

  &::-webkit-scrollbar {
    width: 0.5em;
  }

  .msgBox {
    display: flex;
    flex-direction: column;
    margin: 0;
    padding: 0 1em;
    width: calc(100% - 2em);

    &:last-child {
      margin-bottom: 1em;
    }

    .msg {
      display: flex;
      flex-direction: column;
      width: 100%;
      justify-content: start;
      align-items: start;
      margin-bottom: 0.2em;

      &.myself {
        align-items: end;
        .textBox {
          justify-content: end;
        }
        .text.first {
          border-top-right-radius: 0em;
          border-top-left-radius: 1em;
        }
      }

      &.first {
        margin-top: 1em;

        .text {
          border-top-left-radius: 0em;
        }
      }

      .user {
        display: flex;
        justify-content: start;
        align-items: center;

        .name {
          text-align: start;
          font-size: 0.9em;
          margin-bottom: 0.2em;
        }

        .level {
          text-align: start;
          font-size: x-small;
          color: #74b8f7;
          font-weight: lighter;
          margin: 0 0.25em;
          user-select: none;
        }

        .title {
          border-radius: 0.5em;
          text-align: start;
          font-size: 0.7em;
          height: 80%;
          color: #ffffff;
          font-weight: bold;
          margin: 0 0.25em;
          padding: 0 0.3em;
          user-select: none;
        }
      }

      .textBox {
        display: flex;
      }

      .text {
        width: auto;
        height: auto;
        background-color: #1f1f1f;
        border-radius: 1em;
        font-size: medium;
        /* border-bottom-right-radius: 1em;
        border-bottom-left-radius: 1em; */
        padding: 0.4em 0.8em;

        white-space: normal;
        word-wrap: break-word;
        word-break: break-word;
        text-align: start;
      }

      .time {
        text-align: start;
        display: flex;
        height: 100%;
        align-items: end;
        margin: 0 0.4em;
        margin-bottom: 0.2em;
        font-size: x-small;
        color: #888;
        font-weight: lighter;
        user-select: none;
      }

      img {
        width: 60%;
        border-radius: 1em;
      }

      .file {
        cursor: pointer;
      }
    }
  }
}

.newMsg {
  position: absolute;
  bottom: 3em;
  width: auto;
  height: auto;
  background-color: #1f1f1f;
  border-radius: 1em;
  padding: 0.4em 0.8em;
  color: #888;
  cursor: pointer;

  display: flex;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: center;

  transition: all 0.3s ease-in-out;
  opacity: 0;
  visibility: hidden;

  &:hover {
    background-color: #cecece;
    color: #1f1f1f;
  }

  &.show {
    opacity: 1;
    bottom: 6em;
    visibility: visible;
  }

  .num {
    margin-left: 5px;
    padding: 0 0.2em;
    line-height: 1.5em;

    display: block;
    text-align: center;
    background-color: #74b8f7;
    color: #ffffff;
    border-radius: 0.75em;
    width: auto;
    min-width: 1.1em;
    height: 1.5em;
    text-align: center;

    transition: all 0.3s ease-in-out;
  }
}

.scrollToBtm {
  position: absolute;
  bottom: 3em;
  cursor: pointer;
  transition: all 0.3s ease-in-out;
  transform: rotate(90deg);
  border-radius: 50%;

  width: 3em;
  height: 3em;
  display: flex;
  align-items: center;
  justify-content: center;

  background-color: #74b8f7;

  opacity: 0;
  visibility: hidden;

  &.show {
    opacity: 1;
    bottom: 6em;
    visibility: visible;
  }
}
</style>