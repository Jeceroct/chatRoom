import '@/style/myAlertBox.css'
import MyAlert from './myAlert.js'

const myAlertBox = () => {
  const alertBox = document.createElement('div');
  alertBox.id = 'alertBox';
  return new AlertBox(alertBox);
}

class AlertBox {
  element;
  alerts = [];
  constructor(element) {
    this.element = element;
  }

  showIn(parent) {
    parent.appendChild(this.element);
  }

  add(myAlert) {
    if (myAlert instanceof MyAlert) {
      this.alerts.push(myAlert);
      this.element.appendChild(myAlert.element);
      setTimeout(() => {
        myAlert.show();
      }, 10);
    } else {
      throw new Error('传入的参数必须是是MyAlert类型');
    }
  }

  addClass(className) {
    this.element.classList.add(className);
  }
}

export default myAlertBox;