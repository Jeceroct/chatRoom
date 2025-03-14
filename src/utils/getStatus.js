import request from '../axios'

let instance;

export class Status {

  stop = false
  status = ''
  errorStatus = false

  constructor() {
    if (instance) {
      return instance;
    }
    instance = this;
    this.stop = false;
    this.status = '';
    this.errorStatus = false;
  }

  getStatus(router) {
    if (this.stop) return
    request.post('/getStatus').then(res => {
      switch (res.status) {
        case "checkRoom":
          if (this.status == 'checkRoom') return
          console.log(res)
          this.status = 'checkRoom'
          // lastStatus = res.status
          if (this.errorStatus) {
            this.errorStatus = false
            router.push('/connectionError');
          } else {
            router.push('/address');
          }
          break;
        case "checkUser":
          if (this.status == 'checkUser') return
          console.log(res)
          // lastStatus = res.status
          this.status = 'checkUser'
          router.push('/login');
          break;
        case "start":
          if (this.status == 'start') return
          console.log(res)
          // lastStatus = res.status
          this.status = 'start'
          router.push('/room');
          break;
      }
    }).catch(err => {
      console.log('连接服务器失败', err)
      this.errorStatus = true
      this.status = ''
    })
  }

  stopGetStatus() {
    this.stop = true
    console.log('停止获取状态')
  }

  startGetStatus() {
    this.stop = false
  }

  static getInstance() {
    if (!instance) {
      instance = new Status();
    }
    return instance;
  }
}