<!-- eslint-disable no-unused-vars vue/multi-word-component-names -->
<template>
  <div id="bodyContainer">
    <div class="msgBox"></div>
  </div>
  <div class="newMsg" ref="newMsgRef">
    您有新消息
    <span class="num">{{ newMsgNum }}</span>
  </div>
  <div class="scrollToBtm" ref="scrollToBtmRef">
    <el-icon size="1.8em">
      <DArrowRight />
    </el-icon>
  </div>
</template>

<script setup>
import { onMounted, ref, computed, onUnmounted } from 'vue'
// import { ElMessage } from 'element-plus'
import errorImg from '@/assets/error.jpg'
import avatarError from '@/assets/avatarError.png'
import request from '@/axios'
import axios from "axios";
import myMessage from '@/utils/myMessage'
import { onBeforeRouteLeave } from 'vue-router'
import myMenuItem from '@/utils/myMenuItem';
import { copyToClipboard } from '@/utils/myRightMenuFuncs';
import store from '@/store';

// var user = JSON.parse(localStorage.getItem('chatRoomUserInfo'))
const user = computed(() => {
  const userInfo = JSON.parse(localStorage.getItem('chatRoomUserInfo'))
  return userInfo
})

const newMsgRef = ref(null)
const scrollToBtmRef = ref(null)

const openImgPeriod = ref(true)
const openFilePeriod = ref(true)

const newMsgNum = ref(0)

const downloadingList = new Map([])

const sleep = (delay) => new Promise((resolve) => setTimeout(resolve, delay))

// 防抖函数
let timeoutId;
const debounce = (fn, delay) => {
  return (...args) => {
    clearTimeout(timeoutId);
    timeoutId = setTimeout(() => fn.apply(this, args), delay);
  };
};

const addMsgHTML = async (msg, msgBoxEle) => {
  var newHtml = ''
  if (msgBoxEle.hasChildNodes() && msgBoxEle.lastChild.getAttribute("userId") == msg.from.id) {
    if (msg.from.id == user.value.id) {
      newHtml = `
          <div class="msg myself" userId="${msg.from.id}">
            <div class="textBox">
              <div class="time">${msg.time}</div>
              <div class="text">`+ msg.context + `</div>
            </div>
          </div>`
    } else {
      newHtml = `
            <div class="msg" userId="${msg.from.id}">
              <div class="textBox">
                <div class="text">`+ msg.context + `</div>
                <div class="time">${msg.time}</div>
              </div>
            </div>`
    }
  } else {
    if (msg.from.id == user.value.id) {
      newHtml = `
          <div class="msg first myself" userId="${msg.from.id}">
            <div class="user">
              <div class="title" style="background-color:${msg.from.titleColor}">${msg.from.title}</div>
              <div class="name">${msg.from.name}</div>
              </div>
              <div class="textBox">
                <div class="time">${msg.time}</div>
                <div class="text">`+ msg.context + `</div>
                </div>
                <img class="avatar" onerror="javascript:this.src='${avatarError}'" src="${msg.from.avatar}">
          </div>`
      // <div class="level">Lv${msg.from.level}</div>
    } else {
      newHtml = `
            <div class="msg first" userId="${msg.from.id}">
              <div class="user">
              <img class="avatar" onerror="javascript:this.src='${avatarError}'" src="${msg.from.avatar}">
                <div class="name">${msg.from.name}</div>
                <div class="title" style="background-color:${msg.from.titleColor}">${msg.from.title}</div>
                </div>
                <div class="textBox">
                  <div class="text">`+ msg.context + `</div>
                  <div class="time">${msg.time}</div>
                  </div>
                  </div>`
      // <div class="level">Lv${msg.from.level}</div>
    }
  }
  if (msg.type == 'image') {
    newHtml = newHtml.replace(`<div class="text">` + msg.context + `</div>`, `<img src="${process.env.VUE_APP_API_ADDR}${msg.context}">`)
    msgBoxEle.insertAdjacentHTML('beforeend', newHtml);
    const imgEle = msgBoxEle.lastElementChild.querySelectorAll('img')
    await imgEle.forEach(i => {
      i.onerror = () => {
        i.src = errorImg
      }
      i.style.cursor = 'pointer'
      // 图片左键点击事件
      i.addEventListener('click', () => {
        if (openImgPeriod.value) {
          // 使用新标签页打开
          window.open(process.env.VUE_APP_API_ADDR + msg.context)
          openImgPeriod.value = false
        } else {
          myMessage('图片已打开，请勿多次点击', 'info')
        }
        setTimeout(() => {
          openImgPeriod.value = true
        }, 2000)
      })
      // 图片右键点击事件
      const thisEle = msgBoxEle.lastElementChild
      thisEle.addEventListener('contextmenu', () => {
        store.rightClickMenu.addItem(new myMenuItem('使用 系统图片查看器 打开', () => { request.post('/openImg', msg); myMessage('已在外部打开图片', 'success') }))
        store.rightClickMenu.afterShow(() => { thisEle.classList.add('chosen') })
      })
    })
  } else
    if (msg.type == 'file') {
      newHtml = newHtml.replace(`<div class="text">` + msg.context + `</div>`, `<div class="text file" addr="${msg.context}"><span>文件</span><div class="fileCont">${msg.context}</div></div>`)
      msgBoxEle.insertAdjacentHTML('beforeend', newHtml);
      const fileEle = msgBoxEle.lastElementChild.querySelectorAll('.file')
      await fileEle.forEach(i => {
        i.style.padding = '0'
        i.style.cursor = 'pointer'
        const msgi = {
          title: i.innerText,
          Context: i.getAttribute('addr'),
          From: msg.from
        }
        i.addEventListener('click', () => {
          if (downloadingList.has(msgi.title)) {
            console.log(`${msgi.title}：文件正在下载!`)
            return
          }
          if (openFilePeriod.value) {
            downloadingList.set(msgi.title)
            const elMsg = myMessage(`正在下载文件：${msgi.title}`, 'info', 0)
            request.post('/download', msgi).then(() => {
              elMsg.close()
            }).catch(() => {
              myMessage(`${msgi.title}：下载失败`, 'error')
              elMsg.close()
            }).finally(() => {
              downloadingList.delete(msgi.title)
            })
            openFilePeriod.value = false
          } else {
            myMessage('文件正在下载，请勿多次点击', 'info')
          }
          setTimeout(() => {
            openFilePeriod.value = true
          }, 2000)
        })
        // 文件右键点击事件
        const thisEle = msgBoxEle.lastElementChild
        thisEle.addEventListener('contextmenu', () => {
          // store.rightClickMenu.addItem(new myMenuItem('另存为(未完成)', () => {
          //   const input = document.createElement('input');
          //   input.type = 'file';
          //   // input.setAttribute('webkitdirectory', ''); // 允许选择目录
          //   // input.setAttribute('directory', ''); // 允许选择目录
          //   input.click()
          //   input.onchange = () => {
          //     console.log(input.files)
          //   }
          // }))
          store.rightClickMenu.addItem(new myMenuItem('打开 文件夹', () => {
            request.post('/openDownloadFolder', msgi)
          }))
          store.rightClickMenu.afterShow(() => { thisEle.classList.add('chosen') })
        })
      })
    } else if (msg.type == 'text') {
      msgBoxEle.insertAdjacentHTML('beforeend', newHtml);
      const thisEle = msgBoxEle.lastElementChild
      thisEle.addEventListener('contextmenu', () => {
        store.rightClickMenu.addItem(new myMenuItem('复制 此消息', copyToClipboard, msg.context))
        store.rightClickMenu.afterShow(() => { thisEle.classList.add('chosen') })
      })
    }
}

const initBtnPos = () => {
  const container = document.querySelector('#bodyContainer')
  const scrollToBtmEle = document.querySelector('.scrollToBtm')
  const sendBtnEle = document.querySelector('#sendBtn')
  scrollToBtmEle.style.left = sendBtnEle.offsetLeft + 'px'
  const newMsgElement = document.querySelector('.newMsg')
  newMsgElement.style.left = container.clientWidth / 2 - newMsgElement.clientWidth / 2 + 'px'
}

const isActive = ref(true); // 控制循环是否执行
const abortController = new AbortController(); // 用于取消请求
onMounted(async () => {
  const container = document.querySelector('#bodyContainer')
  // 加载历史消息(只加载最后500条)
  const msgBoxEle = document.querySelector('.msgBox')
  request.post('/getHistory').then(res => {
    if (res) {
      if (res.length > 500) {
        res = res.slice(res.length - 500, res.length)
      }
      console.log(res)
      res.forEach(i => {
        addMsgHTML(i, msgBoxEle)
      });
      container.style.scrollBehavior = 'unset'
      container.scrollTop = container.scrollHeight
      container.style.scrollBehavior = 'smooth'
    }
  })

  // 定义滚动到底部按钮和新消息按钮的位置
  initBtnPos()
  window.addEventListener('resize', debouncedHandler);

  const scrollToBtmEle = document.querySelector('.scrollToBtm')
  const newMsgElement = document.querySelector('.newMsg')
  // 监听container滚动事件，滚动到底部时，隐藏newMsg和scrollToBtm
  container.addEventListener('scroll', debounce(() => {
    const isScrolledToBottom = container.scrollTop + container.clientHeight >= container.scrollHeight - 100
    if (isScrolledToBottom) {
      newMsgElement.classList.remove('show')
      scrollToBtmEle.classList.remove('show')
      newMsgNum.value = 0
    } else if (!newMsgElement.classList.contains('show')) {
      scrollToBtmEle.classList.add('show')
    }
  }), 500)

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
  for (; isActive.value;) {
    const error = ref(null)
    let res
    try {
      res = await request.post('/get', null, { signal: abortController.signal }).catch(err => { error.value = err })
    } catch (err) {
      if (err.name === 'AbortError' || axios.isCancel(err)) {
        console.log('请求已被取消');
        break;
      }
    }
    if (error.value) {
      console.log(error.value)
      await sleep(5000)
    } else {
      console.log(res.message)
      // const isScrolledToBottom = container.scrollTop + container.clientHeight >= container.scrollHeight - 100
      var isScrolledToBottom
      res.message.forEach(i => {
        isScrolledToBottom = !(newMsgRef.value.classList.contains('show') || scrollToBtmRef.value.classList.contains('show'))
        addMsgHTML(i, msgBoxEle)
      });

      // 判断container内部是否已经滚动到最底部
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
  console.log('消息监听循环已停止');
})

const debouncedHandler = debounce(initBtnPos, 250);

onUnmounted(() => {
  isActive.value = false;
  abortController.abort();
  console.log('请求已被取消');
  window.removeEventListener('resize', debouncedHandler);
})

onBeforeRouteLeave((to, from, next) => {
  isActive.value = false;
  abortController.abort();
  console.log('请求已被取消');
  window.removeEventListener('resize', debouncedHandler);
  next()
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
  top: 2.2em;
  padding-bottom: 5px;

  height: calc(100vh - 7em);

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
      width: calc(100% - 2em);
      justify-content: start;
      align-items: start;
      margin-bottom: 0.2em;
      margin-left: 2em;
      position: relative;

      &.myself {
        align-items: end;
        margin-left: 0;
        margin-right: 2em;


        .textBox {
          justify-content: end;
        }

        .avatar {
          left: unset;
          right: -3em;
        }

        &.first .text {
          border-top-right-radius: 0em;
          border-top-left-radius: 1em;
        }
      }

      &.first {
        margin-top: 1.2em;

        .text {
          border-top-left-radius: 0em;
        }
      }

      .user {
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 0.2em;
        height: 1.2em;

        .name {
          display: flex;
          justify-content: center;
          align-items: center;
          font-size: 0.9em;
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
          font-size: 0.7em;
          height: 90%;
          color: #ffffff;
          font-weight: bold;
          margin: 0 0.25em;
          padding: 0 0.3em;
          user-select: none;
          display: flex;
          justify-content: center;
          align-items: center;
        }
      }

      .textBox {
        display: flex;
      }

      .avatar {
        height: 2.5em;
        width: 2.5em;
        border-radius: 50%;
        margin: 0 0.25em;
        user-select: none;
        position: absolute;
        top: 0;
        left: -3em;
      }

      .text {
        width: auto;
        height: auto;
        background-color: var(--color-background);
        color: var(--color-text);
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
        max-width: 60%;
        height: 100px;
        border-radius: 1em;
      }

      .file {
        display: flex;
        flex-direction: column;
        justify-content: center;

        & span {
          font-size: x-small;
          margin: 0.4em 1.3em;
          color: #888;
        }

        & .fileCont {
          margin: 0.5em;
          margin-top: 0;
          padding: 0.4em 0.8em;
          background-color: #ffffff1b;
          border-radius: 1em;
        }
      }

      &.chosen {
        background-color: var(--color-theme);
        padding: 5px;
        border-radius: 10px;
        transform: scale(1.01);
      }

      transition: all 0.2s ease-in-out;
    }
  }
}

.newMsg {
  position: absolute;
  bottom: 3em;
  width: auto;
  height: auto;
  background-color: var(--color-background);
  border-radius: 1em;
  padding: 0.4em 0.8em;
  color: var(--color-text);
  cursor: pointer;

  display: flex;
  flex-wrap: nowrap;
  align-items: center;
  justify-content: center;

  transition: all 0.3s ease-in-out;
  opacity: 0;
  visibility: hidden;

  &:hover {
    background-color: var(--color-hover);
    color: var(--color-hover-text);
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
    background-color: var(--color-theme);
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

  background-color: var(--color-theme);

  opacity: 0;
  visibility: hidden;

  &.show {
    opacity: 1;
    bottom: 6em;
    visibility: visible;
  }
}
</style>