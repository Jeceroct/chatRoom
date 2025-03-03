import request from '../axios'

// var lastStatus = ''
var errorStatus = false
var status = ''

var stop = false

export function stopGetStatus() {
  stop = true
}

export function startGetStatus(router) {
  stop = false
  getStatus(router)
}

function getStatus(router) {
  if (stop) return
  request.post('/getStatus').then(res => {
    switch (res.status) {
      case "checkRoom":
        if (status == 'checkRoom') return
        console.log(res)
        status = 'checkRoom'
        // lastStatus = res.status
        if (errorStatus) {
          errorStatus = false
          router.push('/connectionError');
        } else {
          router.push('/address');
        }
        break;
      case "checkUser":
        if (status == 'checkUser') return
        console.log(res)
        // lastStatus = res.status
        status = 'checkUser'
        router.push('/login');
        break;
      case "start":
        if (status == 'start') return
        console.log(res)
        // lastStatus = res.status
        status = 'start'
        router.push('/room');
        break;
    }
  }).catch(err => {
    console.log('连接服务器失败',err)
    errorStatus = true
    status = ''
  })
}