<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <input type="file" ref="avatarInputEle" accept="image/jpeg,image/png" style="display: none;">
  <input type="file" ref="dataInputEle" accept="text/*" style="display: none;">
  <div class="settingContainer">
    <div class="header">
      <span class="back" @click="hide">&lt;</span>
      <h1>设置</h1>
    </div>
    <div class="userSettings">
      <div class="avatarSetting settings" @click="uploadAvatar">
        <img :src="userInfo.avatar" :onerror="imgError">
        <font-awesome-icon :icon="['fas', 'camera']" />
      </div>
      <div class="nameSetting settings">
        <div class="inputBox">
          <span>昵称</span>
          <input type="text" :value="userInfo.name">
        </div>
      </div>
      <div class="titleSetting settings">
        <div class="inputBox">
          <span>铭牌</span>
          <input type="text" id="titleInput" :value="userInfo.title">
          <div class="color" id="titleColor" @click="titleColorRef.click()"
            :style="`background-color: ${userInfo.titleColor};`">
          </div>
          <input type="color" ref="titleColorRef" :value="userInfo.titleColor" style="display: none;">
        </div>
      </div>
      <button id="titleButton" @click="save">保存</button>
      <div class="moreSettings settings">
        <button id="importData" @click="importData">导入聊天记录</button>
        <button id="exportData" @click="exportData">导出聊天记录</button>
        <button id="logout" @click="logout">退出登录</button>
        <button id="leaveRoom" @click="leaveRoom">退出此聊天室</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import avatarError from '@/assets/avatarError.png'
// import { leave } from '@/utils/leave';
import { onMounted, ref } from 'vue';
import myMessage from '@/utils/myMessage';
import request from '@/axios';
import { Status } from '@/utils/getStatus'
import myDialog from '@/utils/myDialog';

const hide = () => {
  const settingContainer = document.querySelector('.settingContainer');
  settingContainer.classList.remove('show');
}

const userInfo = JSON.parse(localStorage.getItem('chatRoomUserInfo'));

const opacity = ref(0);
const inputColor = ref(userInfo.titleColor)
if (userInfo.titleColor.length !== 7) {
  opacity.value = userInfo.titleColor.slice(-2);
  inputColor.value = userInfo.titleColor.slice(0, 7);
}

const dataInputEle = ref(null);
const avatarInputEle = ref(null);
const titleColorRef = ref(null);

const imgError = (e) => {
  e.target.src = avatarError;
}

const uploadAvatar = () => {
  avatarInputEle.value.click();
  avatarInputEle.value.onchange = (e) => {
    const file = e.target.files[0];
    if (!file) return;
    if (file.type !== 'image/jpeg' && file.type !== 'image/png') {
      myMessage('图片格式错误', 'error');
    }
    const reader = new FileReader();
    reader.readAsDataURL(file)
    reader.onload = (e) => {
      userInfo.avatar = e.target.result;
      save()
    }
  }
}

const save = () => {
  localStorage.setItem('chatRoomUserInfo', JSON.stringify(userInfo));
  request.post('/updateUserInfo', userInfo).then(res => {
    console.log(res);
    if (res.code !== 200) {
      myMessage(res.msg, 'error');
      return;
    }
    myMessage('保存成功', 'success');
    const btn = document.querySelector('#titleButton');
    btn.innerText = '保存成功';
    btn.classList.add('success');
    setTimeout(() => {
      btn.classList.remove('success');
      btn.innerText = '保存';
    }, 3000);
  }).catch(err => {
    console.log(err);
    myMessage('保存失败', 'error');
    const btn = document.querySelector('#titleButton');
    btn.innerText = '保存失败';
    btn.classList.add('error');
    setTimeout(() => {
      btn.classList.remove('error');
      btn.innerText = '保存';
    }, 3000);
  })
}

const importData = () => {
  dataInputEle.value.click();
}

const exportData = () => {
  request.post('/exportData').then(res => {
    // console.log(res);
    // 使用浏览器下载该文件
    const a = document.createElement('a')
    const file = new Blob([JSON.stringify(res)], { type: 'application/json' })
    a.href = URL.createObjectURL(file)
    a.setAttribute('download', '聊天记录.txt')
    a.click()
    myMessage('导出成功', 'success');
  })
}

const logout = () => {
  myDialog('退出登录', 'text', '是否退出登录？', '退出登录', () => {
    localStorage.removeItem('chatRoomUserInfo');
    request.post('/logout').then(res => {
      if (res.code !== 200) {
        myMessage(res.msg, 'error');
        return
      }
      myMessage('退出登录成功', 'success');
    })
  })
}

const leaveRoom = () => {
  myDialog('退出聊天室', 'text', '是否退出聊天室？', '退出聊天室', () => {
    request.post('/leaveRoom').then(res => {
      if (res.code !== 200) {
        myMessage(res.msg, 'error');
        return;
      }
      myMessage('退出成功', 'success');
    })
  })
}

onMounted(() => {
  // 监听铭牌颜色变化
  titleColorRef.value.onchange = (e) => {
    inputColor.value = e.target.value;
    userInfo.titleColor = inputColor.value + opacity.value;
    const titleColorEle = document.querySelector('#titleColor');
    titleColorEle.style.backgroundColor = userInfo.titleColor;
  }

  // 监听导入聊天记录
  dataInputEle.value.onchange = (e) => {
    if (!e.target.files[0]) return;
    if (e.target.files[0].type !== 'text/plain') {
      myMessage('文件格式错误', 'error');
      return;
    }
    const reader = new FileReader();
    reader.readAsText(e.target.files[0]);
    reader.onload = (e) => {
      const data = JSON.parse(e.target.result);
      console.log(data);
      request.post('/importData', reader.result).then(res => {
        if (res.code !== 200) {
          myMessage(res.msg, 'error');
          return;
        }
        myMessage('导入成功', 'success');
      })
    }
  }

  // 监听类名变化
  const observer = new MutationObserver((mutations) => {
    mutations.forEach((mutation) => {
      if (mutation.type === 'attributes' && mutation.attributeName === 'class') {
        const settingContainer = document.querySelector('.settingContainer');
        if (settingContainer.classList.contains('show')) {
          Status.getInstance().startGetStatus()
        } else {
          Status.getInstance().stopGetStatus()
        }
      }
    })
  })
  const targetNode = document.querySelector('.settingContainer');
  const config = { attributes: true };
  observer.observe(targetNode, config);
})
</script>

<style scoped>
.settingContainer {
  width: 100%;
  height: 100%;
  position: absolute;
  top: 0;
  left: 0;
  z-index: 99999;
  background-color: var(--color-background-soft);
  overflow-y: scroll;
  overflow-x: hidden;
  -webkit-overflow-scrolling: touch;
  -ms-overflow-style: none;

  &::-webkit-scrollbar {
    display: none;
  }

  transition: all 0.3s ease;
  transform: translateX(100%);

  &.show {
    transform: translateX(0);
  }
}

.header {
  width: 100%;
  height: 2em;
  padding: 10px;
  background-color: var(--color-background);
  display: flex;
  position: absolute;
  top: 0;
  z-index: 1000;

  .back {
    height: 1.5em;
    width: 1.5em;
    font-size: 1.5em;
    font-weight: bold;
    color: var(--color-text);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all 0.3s ease;
    border-radius: 50%;

    &:hover {
      background-color: var(--color-hover);
      color: var(--color-hover-text);
    }
  }

  h1 {
    font-size: 1.5em;
    font-weight: bold;
    color: var(--color-text);
    margin: 0;
    margin-left: 0.5em;
  }
}

.userSettings {
  width: 100%;
  margin-top: calc(2em + 20px);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: start;
}

.settings {
  width: 100%;
  margin: 0;
}

.moreSettings {
  margin-top: 20px;
}

.avatarSetting {
  width: 8em;
  height: 8em;
  display: block;
  position: relative;
  border-radius: 50%;
  cursor: pointer;
  margin: 2em 0;

  img,
  svg {
    width: 100%;
    height: 100%;
    border-radius: 50%;
    object-fit: cover;
    position: absolute;
    top: 0;
    left: 0;
  }

  img {
    mask-image: linear-gradient(rgba(0, 0, 0, 0.5));
  }

  svg {
    width: 50%;
    height: 50%;
    top: 25%;
    left: 25%;
    opacity: 0.8;
  }
}

button {
  width: 60%;
}

.inputBox {
  display: flex;
  flex-direction: column;
  align-items: start;
  width: 60%;
  margin: 0.3em 20%;

  input {
    width: calc(100% - 2em);
    height: 2.3em;
    padding: 0 1em;
    border-radius: 2em;
    border: 0;
    border-bottom: 5px solid rgb(169, 169, 169);
    transition: all 0.3s ease-in-out;
    margin-bottom: 5px;
  }

  .color {
    width: 100%;
    height: 1.5em;
    border-radius: 2em;
    border: 0;
    transition: all 0.3s ease-in-out;
    margin-bottom: 5px;
    cursor: pointer;
  }

  span {
    opacity: 1;
    color: var(--color-text);
    font-size: 0.8em;
    margin: 0.1em 0 0 0.5em;
    transition: all 0.3s ease-in-out;
  }

  &.error {
    .input {
      border-color: rgb(203, 66, 66);
    }

    span {
      opacity: 1;
    }
  }

  &.hasValue {
    .input {
      border-color: rgb(98, 149, 110);
    }
  }
}
</style>