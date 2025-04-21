<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <routeMask />
  <div class="container">
    <h1>欢迎来到聊天室</h1>
    <el-form @submit.prevent="send">
      <div class="inputBox address">
        <input v-model="addressValue" class="input" placeholder="请输入聊天室的网络地址" />
        <span>此项不能为空</span>
      </div>
      <div class="inputBox password">
        <input v-model="passwordValue" class="input" placeholder="请输入聊天室的密码" />
        <span>警告：此聊天室没有设置密码</span>
      </div>
      <button ref="submitBtn" class="submitBtn" type="submit">确定</button>
    </el-form>

    <div class="alertBoxes"></div>
  </div>
</template>

<script setup>
import routeMask from '../components/routeMask.vue'

import { onMounted, ref } from 'vue'
import request from '../axios'
// import { useRouter } from 'vue-router'
import myAlertBox from '@/utils/myAlertBox';
import MyAlert from '@/utils/myAlert';
// const router = useRouter();

const addressValue = ref('')
const passwordValue = ref('')

const submitBtn = ref(null)

var container

const myAlertBoxEle = myAlertBox()

const send = () => {
  console.log('send')
  addressValue.value = addressValue.value.trim()
  passwordValue.value = passwordValue.value.trim()

  if (addressValue.value === '') {
    const addressEle = document.querySelector('.address')
    addressEle.classList.remove('hasValue')
    addressEle.classList.add('error')
    addressEle.querySelector('input').addEventListener('input', () => {
      addressEle.classList.remove('error')
    })
  }
  if (passwordValue.value === '') {
    const passwordEle = document.querySelector('.password')
    passwordEle.classList.remove('hasValue')
    passwordEle.classList.add('error')
    passwordEle.querySelector('input').addEventListener('input', () => {
      passwordEle.classList.remove('error')
    })
  }

  if (addressValue.value === '') {
    return
  }

  submitBtn.value.setAttribute('disabled', true)
  submitBtn.value.innerHTML = '连接中...'

  const req = new FormData()

  req.append('address', addressValue.value)
  req.append('password', passwordValue.value)
  req.append('db', '0')

  request.post('/address', req, {}).then(res => {
    console.log(res)
    if (res.code == '200') {
      const routeMaskEle = document.querySelector('#routeMask')
      routeMaskEle.classList.add('leave')
      routeMaskEle.classList.remove('waiting')
    } else {
      myAlertBoxEle.add(new MyAlert('聊天室地址或密码错误', 'warning'))
    }
  }).finally(() => {
    submitBtn.value.removeAttribute('disabled')
    submitBtn.value.innerHTML = '确定'
  }).catch(err => {
    console.log(err)
    myAlertBoxEle.add(new MyAlert('聊天室地址或密码错误', 'warning'))
  })
}

onMounted(() => {
  container = document.querySelector('.container')
  myAlertBoxEle.showIn(container)
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

  &[disabled] {
    background-color: rgba(179, 179, 179, 0.496) !important;
    color: #454545 !important;
    cursor: not-allowed !important;
  }
}
</style>