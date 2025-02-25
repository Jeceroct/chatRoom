<template>
  <div id="routeMask">
    加载中...
  </div>
</template>

<script>
export default {
  name: 'routeMask',
  data() {
    return {
      routeMaskEle: null
    }
  },
  mounted() {
    this.routeMaskEle = document.querySelector('#routeMask')
    setTimeout(() => {
      this.routeMaskEle.classList.add('enter')
      setTimeout(() => {
        this.routeMaskEle.classList.add('waiting')
        this.routeMaskEle.classList.remove('enter')
      }, 300)
    }, 100)
  }
}
</script>

<style>
#routeMask {
  position: absolute;
  top: 0;
  right: 0;
  width: 100%;
  height: 100%;
  background-color: #2d2d2d;
  z-index: 99;
  display: flex;
  justify-content: center;
  align-items: center;
  color: white;

  &.enter {
    transition: all 0.1s ease-in-out;
    right: 100%;
  }

  &.waiting {
    right: -100%;
  }

  &.leave {
    transition: all 0.1s ease-in-out;
    right: 0;
  }
}
</style>