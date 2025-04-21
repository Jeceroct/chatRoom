import request from '../axios'
import store from '../store';

const sleep = (delay) => new Promise((resolve) => setTimeout(resolve, delay))

let instance;

export class Status {

  stop = false
  errorStatus = false

  constructor(router) {
    if (instance) {
      return instance;
    }
    instance = this;
    this.stop = false;
    store.currentPage.value = '';
    this.errorStatus = false;
    this.getStatus(router)
  }

  async getStatus(router) {
    for (; ;) {
      await sleep(1000)
      if (this.stop) continue
      if (router == void 0) {
        console.log('router is undefined')
        return
      }
      request.post('/getStatus').then(res => {
        switch (res.status) {
          case "checkRoom":
            if (store.currentPage.value == 'checkRoom') return
            console.log(res)
            store.currentPage.value = 'checkRoom'
            // lastStatus = res.status
            if (this.errorStatus) {
              this.errorStatus = false
              router.push('/connectionError');
            } else {
              router.push('/address');
            }
            break;
          case "checkUser":
            if (store.currentPage.value == 'checkUser') return
            console.log(res)
            // lastStatus = res.status
            store.currentPage.value = 'checkUser'
            router.push('/login');
            break;
          case "start":
            if (store.currentPage.value == 'start') return
            console.log(res)
            // lastStatus = res.status
            store.currentPage.value = 'start'
            router.push('/room');
            break;
        }
      }).catch(err => {
        console.log('连接服务器失败', err)
        this.errorStatus = true
        store.currentPage.value = ''
      })
    }
  }

  stopGetStatus() {
    this.stop = true
    console.log('停止获取状态', this.stop)
  }

  startGetStatus() {
    this.stop = false
    console.log('开始获取状态', this.stop)
  }

  static getInstance(router) {
    if (!instance) {
      instance = new Status(router);
    }
    return instance;
  }
}