import '@/style/myAlert.css'

class MyAlert {
  msg;
  type;
  element;

  constructor(msg, type) {
    this.msg = msg;
    this.type = type;
    this.element = document.createElement('div');
    this.element.className = 'myAlert';
    this.element.classList.add(type);
    this.element.innerHTML = `
      ${this.msg} | ${new Date().toLocaleTimeString()}
    `;
  }

  show() {
    this.element.classList.add('show');
  }

}

export default MyAlert;