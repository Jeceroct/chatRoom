<!-- eslint-disable vue/multi-word-component-names -->
<template>
  <routeMask />
  <div class="container">
    <h1>作为此聊天室的新用户进入</h1>
    <el-form @submit.prevent="send">
      <div class="inputBox id">
        <input v-model="idValue" class="input" placeholder="取一个独特的id" />
        <span>此项不能为空</span>
      </div>
      <div class="inputBox password">
        <input v-model="passwordValue" class="input" placeholder="取一个安全的密码" />
        <span>此项不能为空</span>
      </div>
      <div class="inputBox name">
        <input v-model="nameValue" class="input" placeholder="输入你的昵称" />
        <span>此项不能为空</span>
      </div>
      <div class="inputBox title">
        <input v-model="titleValue" class="input" placeholder="给自己一个头衔吧，不过有没有无所谓了" />
      </div>
      <button id="loginBtn" class="loginBtn" @click="login" type="reset">去登录</button>
      <button id="submitBtn" class="submitBtn" type="submit">注册</button>
    </el-form>
  </div>
</template>

<script setup>
import routeMask from '../components/routeMask.vue'
import { onMounted, ref } from 'vue'
import request from '../axios'
import { useRouter } from 'vue-router'
import { getStatus } from '@/utils/getStatus';
import { ElMessage } from 'element-plus';
const router = useRouter();
// const route = useRoute();
// import { ElMessage } from 'element-plus'

const idValue = ref('')
const passwordValue = ref('')
const nameValue = ref('')
const titleValue = ref('')
const isIdUsed = ref(false)

// 防抖函数
const debounce = (fn, delay) => {
    let timeoutId;
    return (...args) => {
      clearTimeout(timeoutId);
      timeoutId = setTimeout(() => fn.apply(this, args), delay);
    };
  };

const login = () => {
  idValue.value = ''
  const routeMask = document.querySelector('#routeMask')
  routeMask.classList.add('leave')
  routeMask.classList.remove('waiting')
  setTimeout(() => {
    router.push('/login')
  }, 200)
}

const send = () => {
  nameValue.value = nameValue.value.trim()
  passwordValue.value = passwordValue.value.trim()
  idValue.value = idValue.value.trim()
  titleValue.value = titleValue.value.trim()

  if (nameValue.value === '') {
    const nameEle = document.querySelector('.name')
    nameEle.classList.remove('hasValue')
    nameEle.classList.add('error')
    nameEle.querySelector('input').addEventListener('input', () => {
      nameEle.classList.remove('error')
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
  if (idValue.value === '') {
    const idEle = document.querySelector('.id')
    idEle.querySelector('span').innerHTML = '此项不能为空'
    idEle.classList.remove('hasValue')
    idEle.classList.add('error')
    idEle.querySelector('input').addEventListener('input', () => {
      idEle.classList.remove('error')
    })
  }

  if (titleValue.value === '') {
    const titleEle = document.querySelector('.title')
    titleEle.classList.remove('hasValue')
  }

  if (nameValue.value === '' || passwordValue.value === '' || idValue.value === '' || isIdUsed.value) {
    return
  }

  const req = new FormData()

  req.append('id', idValue.value)
  req.append('password', passwordValue.value)
  req.append('name', nameValue.value)
  req.append('title', titleValue.value)

  request.post('/signup', req, {}).then(res => {
    console.log(res)
    if (res.code == 200) {
      const routeMaskEle = document.querySelector('#routeMask')
      routeMaskEle.classList.add('leave')
      routeMaskEle.classList.remove('waiting')
      ElMessage.success('注册成功')
      setTimeout(() => {
        router.push('/login')
      }, 200)
    }
  })
}

onMounted(() => {
  getStatus()
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

  const idEle = document.querySelector('.id .input')
  const errorEle = document.querySelector('.id span')
  const checkIdAvailability = debounce(() => {
    if (!idValue.value.trim()) {
      errorEle.innerHTML = '此项不能为空';
      return;
    }
    request.post('/checkIdUsed', { id: idValue.value }, {}).then(res => {
      if (res.code === '501') {
        errorEle.innerHTML = '此id已被使用';
        idEle.parentElement.classList.remove('hasValue');
        idEle.parentElement.classList.add('error');
        isIdUsed.value = true;
      } else {
        errorEle.innerHTML = '此id可用';
        idEle.parentElement.classList.remove('error');
        isIdUsed.value = false;
      }
    });
  }, 500);
  idEle.addEventListener('input', checkIdAvailability);
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
}
</style>