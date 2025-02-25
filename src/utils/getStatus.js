import request from '../axios'

var status = ''

export function getStatus(router) {
  request.post('/getStatus').then(res => {
    switch (res.status) {
      case "checkRoom":
        if (status == 'checkRoom') return
        console.log(res)
        status = 'checkRoom'
        router.push('/address');
        break;
      case "checkUser":
        if (status == 'checkUser') return
        console.log(res)
        status = 'checkUser'
        router.push('/login');
        break;
      case "start":
        if (status == 'start') return
        console.log(res)
        status = 'start'
        router.push('/room');
        break;
    }
  })
}