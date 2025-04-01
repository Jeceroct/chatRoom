<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <routeMask />
  <div class="container">
    <h1>使用该聊天室的账户登录</h1>
    <el-form @submit.prevent="send">
      <div class="inputBox id">
        <input v-model="idValue" class="input" placeholder="你在此聊天室的用户id" />
        <span>此项不能为空</span>
      </div>
      <div class="inputBox password">
        <input v-model="passwordValue" class="input" placeholder="你在此聊天室的用户的密码" />
        <span>此项不能为空</span>
      </div>
      <button class="loginBtn" @click="signup" type="reset">还没有此聊天室的账户？去注册</button>
      <button ref="submitBtn" class="submitBtn" type="submit">进入</button>
    </el-form>
  </div>
</template>

<script setup>
import routeMask from '../components/routeMask.vue'

import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus';
import request from '../axios'
// import { useRouter } from 'vue-router'
import { leave } from '@/utils/leave';
// const router = useRouter();

const idValue = ref('')
const passwordValue = ref('')

const submitBtn = ref(null)

const signup = () => {
  idValue.value = ''
  leave('/signup')
}

const send = () => {
  idValue.value = idValue.value.trim()
  passwordValue.value = passwordValue.value.trim()
  if (passwordValue.value === '') {
    const passwordEle = document.querySelector('.password')
    passwordEle.classList.remove('hasValue')
    passwordEle.classList.add('error')
    passwordEle.querySelector('input').addEventListener('input', () => {
      passwordEle.classList.remove('error')
    })
  }
  if (idValue.value === '') {
    const idEle = document.querySelector('.id')
    idEle.querySelector('span').innerHTML = '此项不能为空'
    idEle.classList.remove('hasValue')
    idEle.classList.add('error')
    idEle.querySelector('input').addEventListener('input', () => {
      idEle.classList.remove('error')
    })
  }

  if (passwordValue.value === '' || idValue.value === '') {
    return
  }

  submitBtn.value.setAttribute('disabled', true)
  submitBtn.value.innerHTML = '登录中...'

  const req = new FormData()

  req.append('id', idValue.value)
  req.append('password', passwordValue.value)

  request.post('/login', req, {}).then(res => {
    console.log(res)
    if (res.code == '200') {
      const routeMaskEle = document.querySelector('#routeMask')
      routeMaskEle.classList.add('leave')
      routeMaskEle.classList.remove('waiting')
    } else {
      ElMessage.warning('用户Id或密码错误')
    }
  }).finally(() => {
    submitBtn.value.removeAttribute('disabled')
    submitBtn.value.innerHTML = '进入'
  })
}

onMounted(() => {
  const inputs = document.querySelectorAll('.input')
  inputs.forEach((input) => {
    input.addEventListener('input', () => {
      if (input.value !== '') {
        input.parentElement.classList.add('hasValue')
      } else {
        input.parentElement.classList.remove('hasValue')
      }
    })
  })
})
</script>

<style>
.container,
.el-form {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  width: 100%;
}

.inputBox {
  display: flex;
  flex-direction: column;
  align-items: start;
  width: 80%;
  margin: 0.3em 0;

  .input {
    width: calc(100% - 2em);
    height: 2.3em;
    padding: 0 1em;
    border-radius: 2em;
    border: 0;
    border-bottom: 5px solid rgb(169, 169, 169);
    transition: all 0.3s ease-in-out;
  }

  span {
    opacity: 0;
    color: rgb(203, 66, 66);
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

button {
  width: 80%;
  height: 3em;
  border-radius: 2em;
  margin: 0.8em 0;
  background-color: rgb(98, 149, 110);
  border: 2px solid rgb(98, 149, 110);
  color: white;
  cursor: pointer;

  &.loginBtn {
    margin-top: 3em;
    background-color: rgba(214, 255, 224, 0.496);
  }

  &[disabled] {
    background-color: rgba(179, 179, 179, 0.496) !important;
    color: #454545 !important;
    cursor: not-allowed !important;
  }
}
</style>