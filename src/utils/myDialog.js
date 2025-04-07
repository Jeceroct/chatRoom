import '@/style/myDialog.css'

let dialogType = {
  content: 'text',
  image: 'image',
  file: 'file'
}

const myDialog = (title = '标题', contentType = dialogType.content, content = '内容', confirmText = '确定', confirmFunc, ...funcArgs) => {
  const myDialog = document.createElement('div');
  myDialog.classList.add('myDialog');
  myDialog.innerHTML = ''

  if (title != '') {
    myDialog.innerHTML += `
      <div class="myDialog__header">
        <h3>${title}</h3>
      </div>`;
  } else {
    console.warn('myDialog: 标题为空');
  }

  // 判断内容类型
  if (contentType == 'image') {
    myDialog.innerHTML += `
      <div class="myDialog__content">
        <img src="${content}" alt="">
      </div>`
  } else if (contentType == 'file') {
    myDialog.innerHTML += `
      <div class="myDialog__content">
        <div class="myDialog__content__file">${content}</div>
      </div>`
  } else {
    myDialog.innerHTML += `
      <div class="myDialog__content">
        <p>${content}</p>
      </div>`
  }

  myDialog.innerHTML += `
    <div class="myDialog__footer">
      <button class="myDialog__cancel">取消</button>
      <button class="myDialog__confirm">${confirmText}</button>
    </div>`

  document.body.appendChild(myDialog);
  setTimeout(() => {
    myDialog.classList.add('show');
  }, 10);
  return new dialog(myDialog, confirmFunc, ...funcArgs);
}

class dialog {
  element;
  constructor(element, confirmFunc, ...args) {
    this.element = element;
    this.element.querySelector('.myDialog__cancel').addEventListener('click', () => {
      this.close();
    })
    this.element.querySelector('.myDialog__confirm').addEventListener('click', () => {
      this.confirm(confirmFunc, ...args);
    })
  }

  close() {
    this.element.classList.remove('show');
    setTimeout(() => {
      this.element.remove();
    }, 200);
  }

  confirm(confirmFunc, ...args) {
    confirmFunc(...args);
    this.close();
  }
}

export default myDialog;