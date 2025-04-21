import myMessage from "./myMessage"
import myDialog from "./myDialog"
import store from "@/store"
import msg from "@/types/msg"
import RequestType from "@/class/RequestType"
import request from "@/axios"

const copyToClipboard = (text) => {
  navigator.clipboard.writeText(text).then(() => {
    myMessage('复制成功', 'success')
  })
}

const uploadFile = (file) => {
  // 检测是否是图片
  if (file.type.indexOf('image') === -1) {
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = async () => {
      msg.setRequestType(
        RequestType.Type().file,
        RequestType.FileContext(file.name, reader.result),
        `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
        localStorage.getItem('chatRoomUserInfo') ? JSON.parse(localStorage.getItem('chatRoomUserInfo')) : store.user,
        store.quoteValue
      )
      // const elMsg = ElMessage({
      //   message: `正在上传文件: ${file.name}`,
      //   type: 'info',
      //   duration: 0
      // })
      const elMsg = myMessage(`正在上传文件: ${file.name}`, 'info', 0)
      elMsg.load()
      request.post('/send', msg.getResult()).then((res) => {
        if (res.code != 200) {
          myMessage(res.msg, 'error')
          return
        }
        elMsg.close()
      }).catch(() => {
        // ElMessage.error('文件上传失败')
        myMessage('文件上传失败', 'error')
        elMsg.close()
      })
    }
  } else {
    // 获取文件base64内容
    const reader = new FileReader()
    reader.readAsDataURL(file)
    reader.onload = async () => {
      msg.setRequestType(
        RequestType.Type().image,
        reader.result,
        `${new Date().getMonth() + 1}.${new Date().getDate()} ${new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}`,
        localStorage.getItem('chatRoomUserInfo') ? JSON.parse(localStorage.getItem('chatRoomUserInfo')) : store.user,
        store.quoteValue
      )
      // const elMsg = ElMessage({
      //   message: `正在上传图片`,
      //   type: 'info',
      //   duration: 0
      // })
      const elMsg = myMessage(`正在上传图片`, 'info', 0)
      elMsg.load()
      request.post('/send', msg.getResult()).then(() => {
        elMsg.close()
      }).catch(() => {
        elMsg.close()
        // ElMessage.error('图片上传失败')
        myMessage('图片上传失败', 'error')
      })
    }
  }
}

const handlePaste = async (e) => {
  var clipboardData
  var type
  var value
  if (e instanceof ClipboardItem) {
    clipboardData = e
    type = e.types[0]
    console.log(type)
    if (type == void 0) {
      return
    }
    await e.getType(type).then( async (item) => {
      value = await item.text().then((text) => {
        return text
      })
    })
  } else {
    clipboardData = e.clipboardData
    if (clipboardData.files && clipboardData.files.length > 0) {
      type = clipboardData.files[0].type
      value = clipboardData.files[0]
    } else {
      type = clipboardData.types[0]
      value = clipboardData.getData(type)
    }
  }
  // console.log(clipboardData, type, value)
  if (type.indexOf('text') === -1) {
    const reader = new FileReader()
    reader.readAsDataURL(value)
    if (type.indexOf('image') === -1) {
      // 粘贴文件
      reader.onload = () => {
        myDialog('文件', 'file', value.name, '发送文件', uploadFile, value)
      }
    } else {
      // 粘贴图片
      reader.onload = () => {
        myDialog('', 'image', reader.result, '发送图片', uploadFile, value)
      }
    }
  } else {
    // 粘贴的内容不是图片，执行默认粘贴操作
    store.inputValue.value += value
    setTimeout(() => {
      store.inputRef.value.focus()
    }, 0);
  }
}

const downloadFile = () => {

}

const refreshPage = () => {
  window.location.reload()
}

export { copyToClipboard,uploadFile, handlePaste, downloadFile, refreshPage };