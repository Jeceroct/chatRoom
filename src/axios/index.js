import axios from "axios";
// import { ElMessage } from "element-plus";
import router from "@/router/index.js";
import myMessage from "@/utils/myMessage";

const request = axios.create({
  baseURL: process.env.VUE_APP_API_ADDR
  // baseURL: '/api',
  // timeout: 2000
})

var errorShown = false

// request 拦截器
// 可以自请求发送前对请求做一些处理
request.interceptors.request.use(config => {
  config.headers['Content-Type'] = 'application/json;charset=utf-8';
  let user = JSON.parse(localStorage.getItem("xm-user") || '{}')
  config.headers['token'] = user.token || ''
  return config
}, error => {
  // ElMessage.error('无法连接至服务器')
  myMessage('无法连接至服务器', 'error')
  return Promise.reject(error)
});

// response 拦截器
// 可以在接口响应后统一处理结果
request.interceptors.response.use(
  response => {
    let res = response.data;
    // 如果是返回的文件
    if (response.config.responseType === 'blob') {
      return res
    }
    // 当权限验证不通过的时候给出提示
    if (res.code === '401') {
      // ElMessage.error(res.msg)
      myMessage(res.msg, 'error')
      router.push('/login')
    }
    // 兼容服务端返回的字符串数据
    // if (typeof res === 'string') {
    //   res = res ? JSON.parse(res) : res
    // }
    return res;
  },
  error => {
    if (errorShown) {
      return Promise.reject(error)
    }
    errorShown = true
    if (error.response.status === 404) {
      // ElMessage.error('未知的请求接口')
      myMessage('未知的请求接口', 'error')
    } else if (error.response.status === 500) {
      // ElMessage.error('无法连接至服务器')
      myMessage('无法连接至服务器', 'error')
    } else if (error.response.status === 501) {
      // ElMessage.error('Redis连接失败')
      myMessage('Redis连接失败', 'error')
    } else if (error.response.status === 901) {
      // ElMessage.error('未知的消息格式')
      myMessage('未知的消息格式', 'error')
    } else if (error.response.status === 902) {
      // ElMessage.error('文件发送失败')
      myMessage('文件发送失败', 'error')
    } else if (error.response.status === 903) {
      // ElMessage.error('文件接收失败')
      myMessage('文件接收失败', 'error')
    } else if (error.response.status === 904) {
      // ElMessage.error('文件创建失败')
      myMessage('文件创建失败', 'error')
    } else if (error.response.status === 905) {
      // ElMessage.error('未知的用户信息格式')
      myMessage('未知的用户信息格式', 'error')
    } else if (error.response.status === 906) {
      // ElMessage.error('用户信息不存在')
      myMessage('用户信息不存在', 'error')
    } else {
      // ElMessage.error('未知错误')
      myMessage('未知错误', 'error')
      console.error(error.message)
    }
    setTimeout(() => {
      errorShown = false
    }, 4000)
    console.warn(error.response)
    return Promise.reject(error)
  }
)

export default request
