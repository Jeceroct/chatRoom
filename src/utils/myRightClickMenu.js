import myMenuItem from './myMenuItem.js';
import '@/style/myRightClickMenu.css'

const initRightClickMenu = (menu) => {
  const menuEle = document.createElement('div');
  menuEle.classList.add('right-click-menu');
  menuEle.classList.add('hide');
  menu.items.forEach(item => {
    const itemEle = document.createElement('div');
    itemEle.classList.add('right-click-menu-item');
    itemEle.innerText = item.getTitle();
    itemEle.addEventListener('click', () => {
      item.submit();
    })
    menuEle.appendChild(itemEle);
  })
  return menuEle;
}

const showRightClickMenu = (menuEle, x, y) => {
  document.body.appendChild(menuEle);
  if (x + menuEle.offsetWidth > window.innerWidth) x = x - menuEle.offsetWidth;
  if (y + menuEle.offsetHeight > window.innerHeight) y = y - menuEle.offsetHeight;
  menuEle.style.left = `${x}px`;
  menuEle.style.top = `${y}px`;
  setTimeout(() => {
    menuEle.classList.remove('hide');
  }, 0);
}

const closeRightClickMenu = () => {
  // 关闭所有的右键菜单
  const elements = document.querySelectorAll('.right-click-menu');
  elements.forEach(element => {
    element.classList.add('hide');
    setTimeout(() => {
      element.remove();
    }, 500);
  })

  // 去除所有chosen样式
  const msgs = document.querySelectorAll('.chosen');
  msgs.forEach(msg => {
    msg.classList.remove('chosen');
  })
}

class myRightClickMenu {
  items;
  element;
  afterShowCallbacks = [];

  constructor(...items) {
    closeRightClickMenu();
    if (!items.every(item => item instanceof myMenuItem)) throw new Error('右键菜单的选项必须是myMenuItem类型')
    this.items = items;
    this.element = initRightClickMenu(this);
    return this;
  }

  show(posX, posY) {
    this.element = initRightClickMenu(this);
    closeRightClickMenu();
    showRightClickMenu(this.element, posX, posY);
    for (const callback of this.afterShowCallbacks) {
      callback();
    }
    this.afterShowCallbacks = [];
  }

  static close() {
    closeRightClickMenu();
  }

  addItem(item) {
    if (!(item instanceof myMenuItem)) throw new Error('右键菜单的选项必须是myMenuItem类型')
    // this.items.unshift(item);
    this.items.push(item);
    this.element = initRightClickMenu(this);
  }

  removeItem(item) {
    this.items = this.items.filter(i => i !== item);
    this.element = initRightClickMenu(this);
  }

  removeAllItems() {
    this.items = [];
    this.element = initRightClickMenu(this);
  }

  afterShow(callback, ...args) {
    if (typeof callback !== 'function') {
      throw new Error('传入的 callback 必须是一个函数');
    }
    this.afterShowCallbacks.push(() => {
      callback(...args);
    })
  }
}

export default myRightClickMenu;