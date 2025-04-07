import '@/style/myMessage.css'

let index = 0;

const messageType = {
  info: 'info',
  success: 'success',
  warning: 'warning',
  error: 'error'
}

const myMessage = (content, type = messageType.info, duration = 3000) => {
  const myMessage = document.createElement('div');
  myMessage.classList.add('myMessage');
  myMessage.classList.add(type);
  myMessage.innerHTML = `
        <div class="myMessage__content">
            <p>${content}</p>
        </div>
    `;
  document.body.appendChild(myMessage);
  index++;
  myMessage.style.setProperty('--index', index);
  setTimeout(() => {
    myMessage.classList.add('myMessage__show');
    if (duration === 0) {
      return new message(myMessage);
    }
    setTimeout(() => {
      close(myMessage);
    }, duration)
  }, 50);
  return new message(myMessage);
}

class message {
  content

  constructor(myMessage) {
    this.content = myMessage;
  }

  close() {
    close(this.content);
  }

  load() {
    const loading = document.createElement('div');
    loading.classList.add('myMessage__loading');
    this.content.appendChild(loading);
  }
}

const close = (myMessage) => {
  index--;
  myMessage.classList.remove('myMessage__show');
  document.querySelectorAll('.myMessage').forEach((item) => {
    item.style.setProperty('--index', item.style.getPropertyValue('--index') - 1);
  })
  setTimeout(() => {
    myMessage.remove();
  }, 200);
}

export default myMessage;
export { messageType }