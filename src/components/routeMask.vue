<template>
  <div id="routeMask">
    加载中...
  </div>
</template>

<script setup>
import { onMounted,defineExpose } from 'vue'

var routeMaskEle

defineExpose({
  routeMaskEle
})

onMounted(() =>{
  routeMaskEle = document.querySelector('#routeMask')
  setTimeout(() => {
    routeMaskEle.classList.add('enter')
    setTimeout(() => {
      routeMaskEle.classList.add('waiting')
      routeMaskEle.classList.remove('enter')
    }, 300)
  }, 100)
})
</script>

<style>
#routeMask {
  position: absolute;
  top: 0;
  right: 0;
  width: 100%;
  height: 100%;
  background-color: var(--color-background);
  z-index: 9999;
  display: flex;
  justify-content: center;
  align-items: center;
  color: var(--color-text);

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